package table

type Version string
type ID string
type ColumnName string

type TypedCell struct {
	Int64  *int64
	String *string
}

type TypedRow map[ColumnName]TypedCell

type Row map[ColumnName]string
type Rows []Row

type Table struct {
	ID         ID       `json:"id"`
	Version    Version  `json:"version"`
	Timestamp  int64    `json:"timestamp"`
	Rows       Rows     `json:"rows"`
	ReadErrors []string `json:"readErrors"`
}

type TableSummary struct {
	ID        ID      `json:"id"`
	Timestamp int64   `json:"timestamp"`
	Version   Version `json:"version"`
	Valid     bool    `json:"valid"`
	Committed bool    `json:"committed"`
}

type List []TableSummary
type Map map[ID]List

func (l List) Len() int {
	return len(l)
}

func (l List) Less(i, j int) bool {
	return l[i].Timestamp < l[j].Timestamp
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (t *Table) AppendRow(row Row, lineError error) {
	t.Rows = append(t.Rows, row)
	e := ""
	if lineError != nil {
		e = lineError.Error()
	}
	t.ReadErrors = append(t.ReadErrors, e)
}
