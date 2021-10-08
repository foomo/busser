package service

import (
	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
)

type Service interface {
	Validate(id table.ID) (t *table.Table, vt *validation.Table, err *ErrorValidation)
	GetVersion(id table.ID, version table.Version) (t *table.Table, vt *validation.Table, err *ErrorGet)
	GetCommitted(id table.ID) (t *table.Table, vt *validation.Table, err *ErrorGet)
	Delete(id table.ID, versions []table.Version) (err *ErrorDelete)
	Commit(id table.ID, version table.Version) *ErrorCommit
	List() (table.Map, *ErrorGet)
}
