package processor

import (
	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
)

// Processor is an interface, that allows you to validate and transform table data typically from a csv file.
type Processor interface {
	// called for every cell first in order
	Cell(
		collector validation.Collector,
		col table.ColumnName,
		content string,
	) (
		cleanedContent string, valid bool, err error,
	)
	// called after all cells of all rows were processed
	Row(
		collector validation.Collector,
		row table.Row,
	) (valid bool, err error)
	// called in the very end
	Table(
		collector validation.Collector,
		tableValidation *validation.Table,
		table *table.Table,
	) (valid bool, err error)
}

func Process(t *table.Table, p Processor) (vt *validation.Table, err error) {
	vt = &validation.Table{
		Valid: false,
	}
	vt.Rows = make(validation.Rows, len(t.Rows))
	c := &validation.Container{}
	for i, row := range t.Rows {
		rowValidation := &validation.Row{}
		vt.Rows[i] = rowValidation
		rowValidation.Cells = make(map[table.ColumnName]validation.Cell)
		for colName, cellContent := range row {
			cleanContent, valid, err := p.Cell(
				c.Collect,
				colName,
				cellContent,
			)
			rowValidation.Cells[colName] = validation.Cell{
				Valid:    valid,
				Feedback: c.Flush(),
			}
			if err != nil {
				return nil, err
			}
			row[colName] = cleanContent
		}
	}
	for i, row := range t.Rows {
		rv := vt.Rows[i]
		// do validation on row
		valid, err := p.Row(
			c.Collect,
			row,
		)
		if err != nil {
			return nil, err
		}
		rv.Valid = valid

		// if the row validations are valid check if
		// the cell validations are valid as well and
		// propagate the validation state
		if rv.Valid {
			rv.Valid = rv.CellsAreValid()
		}

		rv.Feedback = c.Flush()
	}
	c.Feedback = nil
	valid, err := p.Table(c.Collect,
		vt,
		t,
	)

	// if the table validations are valid check if
	// the row validations are valid as well and
	// propagate the validation state
	if valid {
		valid = vt.RowsAreValid()
	}
	vt.Valid = valid
	vt.Feedback = c.Flush()
	return vt, err
}
