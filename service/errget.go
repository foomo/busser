package service

type ErrorGet string

const (
	ErrorGetCouldNotLoadTableFromStore ErrorGet = "COULD_NOT_LOAD_TABLE_FROM_STORE"
)

var (
	ErrCouldNotLoadTableFromStore = newGetError(ErrorGetCouldNotLoadTableFromStore)
)

func newGetError(e ErrorGet) *ErrorGet {
	return &e
}

func (e *ErrorGet) Error() string {
	return string(*e)
}
