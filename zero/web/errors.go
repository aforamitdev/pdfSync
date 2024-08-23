package web

import "errors"

type ErrorResponse struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// passed when request is invalid
type RequestError struct {
	Err    error
	Status int
}

// Error implements error.
func (re *RequestError) Error() string {
	return re.Err.Error()
}

func NewRequestError(err error, status int) error {
	return &RequestError{err, status}
}

func IsRequestError(err error) bool {
	var re *RequestError
	return errors.As(err, &re)
}

func GetRequestError(err error) *RequestError {
	var re *RequestError
	if !errors.As(err, &re) {
		return nil
	}
	return re
}
