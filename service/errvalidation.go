package service

type ErrorValidation string

const (
	ErrorValidationCouldNotValidate ErrorValidation = "COULD_NOT_LOAD_TABLE_FROM_STORE"
)

var (
	ErrValidationCouldNotValidate = newValidationError(ErrorValidationCouldNotValidate)
)

func newValidationError(e ErrorValidation) *ErrorValidation {
	return &e
}

func (e *ErrorValidation) Error() string {
	return string(*e)
}
