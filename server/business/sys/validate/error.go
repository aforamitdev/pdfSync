package validate

import (
	"encoding/json"
	"errors"
)

var ErrInvalidID = errors.New("ID is not in its proper form")

type ErrorResponse struct {
	Error string `json:"error"`
	Field string `json:"fields,omitempty"`
}

type RequestError struct {
	Err    error
	status int
	fields error
}

func NewRequestError(err error, status int) error {
	return &RequestError{Err: err, status: status, nil}
}

func (err *RequestError) Error() string {
	return err.Err.Error()
}

type FieldError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type FieldErrors []FieldError

func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}
