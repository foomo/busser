package validation

import "github.com/foomo/busser/table"

type FeedbackLevel string

const (
	FeedbackLevelValid   FeedbackLevel = "valid"
	FeedbackLevelWarning FeedbackLevel = "warning"
	FeedbackLevelError   FeedbackLevel = "error"
)

type FeedbackEntry struct {
	Level FeedbackLevel `json:"level"`
	Msg   string        `json:"msg"`
}

type Feedback []FeedbackEntry

type Cell struct {
	Valid    bool     `json:"valid"`
	Feedback Feedback `json:"feedback"`
}

type Row struct {
	Valid    bool                      `json:"valid"`
	Cells    map[table.ColumnName]Cell `json:"cells"`
	Feedback Feedback                  `json:"feedback"`
}

func (r *Row) AddFeedback(level FeedbackLevel, msg string) {
	r.Feedback = append(r.Feedback, FeedbackEntry{Level: level, Msg: msg})
}

func (r Row) CellsAreValid() bool {
	for _, c := range r.Cells {
		if !c.Valid {
			return false
		}
	}
	return true
}

type Rows []*Row

type Table struct {
	Valid    bool     `json:"valid"`
	Rows     Rows     `json:"rows"`
	Feedback Feedback `json:"feedback"`
}

func (t Table) RowsAreValid() bool {
	for _, r := range t.Rows {
		if !r.Valid {
			return false
		}
	}
	return true
}

type Collector func(level FeedbackLevel, msg string)

type Container struct {
	Feedback `json:"feedback"`
}

func (vc *Container) Collect(level FeedbackLevel, msg string) {
	vc.Feedback = append(vc.Feedback, FeedbackEntry{Level: level, Msg: msg})
}

func (vc *Container) Flush() Feedback {
	feedback := vc.Feedback
	vc.Feedback = nil
	return feedback
}
