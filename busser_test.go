package busser

import (
	"context"
	"errors"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/foomo/busser/config"
	"github.com/foomo/busser/csvreader"
	"github.com/foomo/busser/service"
	"github.com/foomo/busser/store"
	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type exampleProcessor struct{}

// called for every cell first in order
func (p *exampleProcessor) Cell(
	c validation.Collector,
	col table.ColumnName,
	content string,
) (
	cleanedContent string, valid bool, err error,
) {
	cleanedContent = strings.Trim(content, " 	")
	valid = true
	switch col {
	case "From", "To":
		if cleanedContent == "" {
			c(validation.FeedbackLevelError, "must not be empty")
			valid = false
		}

	case "Enabled":
		switch cleanedContent {
		case "yes", "no":
		default:
			valid = false
			c(validation.FeedbackLevelError, "must be yes or no")
		}
	}
	return cleanedContent, valid, nil
}

func (p *exampleProcessor) Row(
	c validation.Collector,
	row table.Row,
) (valid bool, err error) {
	if row["From"] == row["To"] {
		c(validation.FeedbackLevelError, "from and to must differ")
		return false, nil
	}
	return true, nil
}

func (p *exampleProcessor) Table(
	collector validation.Collector,
	tableValidation *validation.Table,
	table *table.Table,
) (valid bool, err error) {
	dupls := map[string]int{}
	for i, row := range table.Rows {
		duplKey := row["From"] + "-" + row["To"]
		_, isDupl := dupls[duplKey]
		if isDupl {
			tableValidation.Rows[i].Valid = false
			collector(validation.FeedbackLevelError, "duplicate entry found")
			tableValidation.Rows[i].AddFeedback(validation.FeedbackLevelError, "duplicate line")
		}
		dupls[duplKey]++
	}
	return true, nil
}

const exampleID table.ID = "example"

func getTestData(t *testing.T) (
	ctx context.Context,
	b *Busser,
	client *service.HTTPServiceGoTSRPCClient,
) {
	dir, err := os.MkdirTemp("", "store-test")
	assert.NoError(t, err)
	s, err := store.NewFS(dir)
	assert.NoError(t, err)
	l, err := zap.NewProduction()
	assert.NoError(t, err)
	defer l.Sync()
	ctx = context.Background()
	b = New(ctx, l, s,
		config.Config{
			exampleID: config.Table{
				Processor: &exampleProcessor{},
				Loader:    csvreader.GetByteTableLoader(exampleID, []byte(testTable), nil),
			},
		},
	)
	p := service.NewDefaultServiceGoTSRPCProxy(b)
	server := httptest.NewServer(p)
	client = service.NewDefaultServiceGoTSRPCClient(server.URL)
	return ctx, b, client
}

const testTable = `From,To,Enabled,Comment
/from,/to,yes,this is a valid test
/foo,/bar,,this is missing Enabled
/foo,,no,this is missing to
/same,/same,yes,"this is invalid, because From and To are the same"
/foo,/bar,yes,this is a duplicate`

func Test(t *testing.T) {
	ctx, _, client := getTestData(t)
	tableMap, errList, errClient := client.List(ctx)
	if errList != nil {
		t.Fatal("list error", errList)
		return
	}
	assert.NoError(t, errClient)
	tbl, vt, errValidate, _ := client.Validate(ctx, exampleID)
	if errValidate != nil {
		t.Fatal("unexpected validation error", errValidate)
		return
	}
	assert.False(t, vt.Rows[3].Valid)
	assert.Len(t, tbl.Rows, 5)
	tableMap, errList, errClient = client.List(ctx)
	assert.NoError(t, errClient)
	if errList != nil {
		t.Fatal("list error", errList)
		return
	}
	assert.Len(t, tableMap[exampleID], 1)
	assert.Equal(t, tbl.Version, tableMap[exampleID][0].Version)
	errCommit, errClient := client.Commit(ctx, tbl.ID, tbl.Version)
	assert.NoError(t, errClient)
	if errCommit != nil {
		t.Fatal("could not commit table", errCommit)
		return
	}

	tbl, vt, errGet, _ := client.GetCommitted(ctx, exampleID)
	if errGet != nil {
		t.Fatal("could not get committed version", errGet)
		return
	}

	assert.True(t, vt.Rows[0].Valid)
	assert.False(t, vt.Rows[1].Valid)
	assert.False(t, vt.Rows[3].Valid)
	assert.False(t, vt.Rows[4].Valid)
	assert.Len(t, vt.Rows[4].Feedback, 1)
	assert.Equal(t, vt.Rows[4].Feedback[0].Msg, "duplicate line")
	assert.Len(t, vt.Feedback, 1)
	assert.Equal(t, vt.Feedback[0].Msg, "duplicate entry found")

	errDelete, errClient := client.Delete(ctx, exampleID, []table.Version{tbl.Version})
	assert.NoError(t, errClient)
	if !errors.Is(errDelete, service.ErrDeleteCouldNotDeleteVersion) {
		t.Fatal("committed table version must be protected", errDelete)
		return
	}
	nextTbl, _, errValidate, errClient := client.Validate(ctx, exampleID)
	assert.NoError(t, errClient)

	if errValidate != nil {
		t.Fatal("could not validate table", errValidate)
		return
	}

	errCommit, errClient = client.Commit(ctx, nextTbl.ID, nextTbl.Version)
	assert.NoError(t, errClient)
	if errCommit != nil {
		t.Fatal("could not commit table", errCommit)
		return
	}

	errDelete, errClient = client.Delete(ctx, exampleID, []table.Version{tbl.Version})
	assert.NoError(t, errClient)
	if errDelete != nil {
		t.Fatal("could not delete previously committed table", errDelete)
		return
	}

	tableMap, errList, errClient = client.List(ctx)
	assert.NoError(t, errClient)
	if errList != nil {
		t.Fatal("list error", errList)
		return
	}
	assert.Len(t, tableMap[exampleID], 1)

}
