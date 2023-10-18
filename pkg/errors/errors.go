package errors

type Error struct {
	Code    int
	Message string
	Err     error
}

// New create error
func New(code int, msg string) *Error {
	return &Error{Code: code, Message: msg}
}

// Error return error with info
func (e *Error) Error() string {
	return e.Message
}

// WithMsg with message
func (e *Error) WithMsg(message string) *Error {
	e.Message = message
	return e
}

// WithError with original error
func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}
