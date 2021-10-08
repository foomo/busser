package service

type ErrorDelete string

const (
	ErrorDeleteCouldNotDeleteVersion ErrorDelete = "COULD_NOT_DELETE_VERSION"
)

var (
	ErrDeleteCouldNotDeleteVersion = newDeleteError(ErrorDeleteCouldNotDeleteVersion)
)

func newDeleteError(e ErrorDelete) *ErrorDelete {
	return &e
}

func (e *ErrorDelete) Error() string {
	return string(*e)
}

func (e *ErrorDelete) Is(err error) bool {
	switch ee := err.(type) {
	case *ErrorDelete:
		return ee.Error() == e.Error()
	}
	return false
}
