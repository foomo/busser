package config

import (
	"github.com/foomo/busser/processor"
	"github.com/foomo/busser/table"
)

type TableLoader func() (t *table.Table, err error)

type Table struct {
	ID        table.ID
	Processor processor.Processor
	Loader    TableLoader
}

type Config map[table.ID]Table
