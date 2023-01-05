package apperror

const (
	NotFoundError = iota + 1
	AuthError
)

type CommonError struct {
	ParentError error `json:"-"`
	Code        int
	Message     string
}

func (e CommonError) Error() string {
	return e.Message
}

func NewError(err error, message string, code int) error {
	return CommonError{ParentError: err, Message: message, Code: code}
}
