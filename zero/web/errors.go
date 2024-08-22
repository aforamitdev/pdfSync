package web

import "github.com/pkg/errors"

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ErrorResponse struct {
	Error  string       `json:"error"`
	Fields []FieldError `json:"fields,omitempty"`
}

// RequestError is used to pass an error during the request through the
// application with web specific context.
type RequestError struct {
	Err    error
	Status int
}

func (re *RequestError) Error() string {
	return re.Err.Error()
}

type Error struct {
	Err    error
	Status int
	Fields []FieldError
}

// quelifys the error as error
func (err *Error) Error() string {
	return err.Err.Error()
}

func NewRequestError(err error, status int) error {
	return &Error{err, status, []FieldError{}}
}

type shutdown struct {
	Message string
}

func (s *shutdown) Error() string {
	return s.Message
}

func IsShutdown(err error) bool {
	if _, ok := errors.Cause(err).(*shutdown); ok {
		return true
	}
	return false
}

func GetRequestError(err error) *RequestError {
	var re *RequestError
	if !errors.As(err, &re) {
		return re
	}
	return re
}
