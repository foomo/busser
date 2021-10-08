package store

import (
	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
)

type Store interface {
	Add(t *table.Table, vt *validation.Table) error
	GetVersion(id table.ID, version table.Version) (t *table.Table, vt *validation.Table, err error)
	List() (table.List, error)
	Delete(id table.ID, version table.Version) error
	Commit(id table.ID, version table.Version) error
	GetCommitted(id table.ID) (t *table.Table, vt *validation.Table, err error)
}
