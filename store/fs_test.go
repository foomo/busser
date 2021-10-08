package store

import (
	"io/ioutil"
	"testing"

	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	tbl := &table.Table{
		ID:      "test-table",
		Version: "one",
		Rows: table.Rows{table.Row{
			"foo": "bar",
		}},
	}
	vt := &validation.Table{
		Valid: true,
	}
	testRoot, err := ioutil.TempDir("", "busser-fs-test-store-")
	assert.NoError(t, err)
	fs, err := NewFS(testRoot)
	assert.NoError(t, err)
	assert.NoError(t, fs.Add(tbl, vt))
	list, err := fs.List()
	assert.NoError(t, err)
	assert.Len(t, list, 1)
	assert.NoError(t, fs.Delete(tbl.ID, tbl.Version))
	list, err = fs.List()
	assert.NoError(t, err)
	assert.Len(t, list, 0)

}
