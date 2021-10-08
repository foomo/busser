package csvreader

import (
	"bytes"
	"testing"

	"github.com/foomo/busser/table"

	"github.com/stretchr/testify/assert"
)

const testCSV = `A,B,C
/foo,,
X
a,b,c
`

func Test(t *testing.T) {
	tbl, err := Read(bytes.NewBuffer([]byte(testCSV)), nil)
	assert.NoError(t, err)
	assert.Len(t, tbl.Rows, 3)
	assert.Equal(t, "X", tbl.Rows[1]["A"])
	assert.Equal(t, table.Row{"A": "X", "B": "", "C": ""}, tbl.Rows[1])
	assert.Equal(t, table.Row{"A": "a", "B": "b", "C": "c"}, tbl.Rows[2])
	assert.Equal(t, table.Row{"A": "/foo", "B": "", "C": ""}, tbl.Rows[0])
}
