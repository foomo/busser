package busser

import (
	"context"
	"sort"

	"github.com/foomo/busser/config"
	"github.com/foomo/busser/processor"
	"github.com/foomo/busser/service"
	"github.com/foomo/busser/store"
	"github.com/foomo/busser/table"
	"github.com/foomo/busser/table/validation"
	"go.uber.org/zap"
)

type Busser struct {
	store store.Store
	conf  config.Config
	l     *zap.Logger
	ctx   context.Context
}

func New(
	ctx context.Context,
	l *zap.Logger,
	store store.Store,
	conf config.Config,
) *Busser {
	return &Busser{
		ctx:   ctx,
		l:     l,
		store: store,
		conf:  conf,
	}
}

func (b *Busser) ml(method string) *zap.Logger {
	return b.l.With(zap.String("method", method))
}

func (b *Busser) Validate(id table.ID) (
	t *table.Table,
	vt *validation.Table,
	err *service.ErrorValidation,
) {
	l := b.ml("validate-table")
	l.Info("validating", zap.String("table", string(id)))
	cnf, ok := b.conf[id]
	if !ok {
		l.Error("conf not found")
		return nil, nil, service.ErrValidationCouldNotValidate
	}
	t, errLoad := cnf.Loader()
	if errLoad != nil {
		l.Error("loading of table failed", zap.Error(errLoad))
		return nil, nil, service.ErrValidationCouldNotValidate
	}
	if t.ID == "" {
		l.Error("loader returned table without id")
		return nil, nil, service.ErrValidationCouldNotValidate
	}
	if t.Version == "" {
		l.Error("loader returned table without version")
		return nil, nil, service.ErrValidationCouldNotValidate
	}

	vt, errProcess := processor.Process(t, cnf.Processor)
	if errProcess != nil {
		l.Error("processing failed", zap.Error(errProcess))
		return nil, nil, service.ErrValidationCouldNotValidate
	}
	errAdd := b.store.Add(t, vt)
	if errAdd != nil {
		l.Error("could not add table to store", zap.Error(errAdd))
		return nil, nil, service.ErrValidationCouldNotValidate
	}
	l.Info("validation complete, added version to store", zap.String("table-version", string(t.Version)), zap.String("table-id", string(t.ID)))
	return t, vt, nil
}
func (b *Busser) GetVersion(id table.ID, version table.Version) (t *table.Table, vt *validation.Table, err *service.ErrorGet) {
	t, vt, errGetVersion := b.store.GetVersion(id, version)
	if errGetVersion != nil {
		b.ml("get-version").Error("could not get version from store", zap.Error(errGetVersion))
		return nil, nil, service.ErrCouldNotLoadTableFromStore
	}
	return t, vt, nil
}
func (b *Busser) Commit(id table.ID, version table.Version) *service.ErrorCommit {
	errCommit := b.store.Commit(id, version)
	if errCommit != nil {
		b.ml("commit-table").Error("could not commit table", zap.Error(errCommit))
		return service.ErrCommitCouldNotCommit
	}
	return nil
}

func (b *Busser) GetCommitted(id table.ID) (t *table.Table, vt *validation.Table, err *service.ErrorGet) {
	t, vt, errGetCommitted := b.store.GetCommitted(id)
	if errGetCommitted != nil {
		b.ml("get-commited").Error("could not get committed table", zap.Error(errGetCommitted))
		return nil, nil, service.ErrCouldNotLoadTableFromStore
	}
	return t, vt, nil
}

func (b *Busser) Delete(id table.ID, versions []table.Version) *service.ErrorDelete {
	l := b.ml("delete").With(zap.String("table-id", string(id)))
	committedTable, _, errGetCommitted := b.store.GetCommitted(id)
	if errGetCommitted != nil {
		l.Error("could not determine committed version", zap.Error(errGetCommitted))
		return service.ErrDeleteCouldNotDeleteVersion
	}
	for _, v := range versions {
		zv := zap.String("table-version", string(v))
		if committedTable != nil && committedTable.Version == v {
			l.Error("can not delete commited version")
			return service.ErrDeleteCouldNotDeleteVersion
		}
		err := b.store.Delete(id, v)
		if err != nil {
			l.Error("could not delete version", zv)
			return service.ErrDeleteCouldNotDeleteVersion
		}
		l.Info("delete table version", zv)
	}
	return nil
}

func (b *Busser) List() (table.Map, *service.ErrorGet) {
	l := b.ml("list")
	list, err := b.store.List()
	if err != nil {
		l.Error("could not list", zap.Error(err))
		return nil, service.ErrCouldNotLoadTableFromStore
	}
	tableMap := table.Map{}
	for tableID := range b.conf {
		tableMap[tableID] = table.List{}
		for _, ts := range list {
			if ts.ID == tableID {
				tableMap[tableID] = append(tableMap[tableID], ts)
			}
		}
		sort.Sort(tableMap[tableID])
	}
	return tableMap, nil
}
