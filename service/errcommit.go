package service

type ErrorCommit string

const (
	ErrorCommitCouldNotCommit ErrorCommit = "COULD_NOT_COMMIT"
)

var (
	ErrCommitCouldNotCommit = newCommitError(ErrorCommitCouldNotCommit)
)

func newCommitError(e ErrorCommit) *ErrorCommit {
	return &e
}

func (e *ErrorCommit) Error() string {
	return string(*e)
}
