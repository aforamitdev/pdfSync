package validate

import (
	"encoding/json"
	"errors"
	"fmt"
)

type FieldError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

type FieldErrors []FieldError

// NewFieldsError creates an fields error.
func NewFieldsError(field string, err error) error {
	return FieldErrors{
		{
			Field: field,
			Err:   err.Error(),
		},
	}
}

func (fe FieldErrors) Error() string {
	d, err := json.Marshal(fe)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

func (fe FieldErrors) Fields() map[string]string {
	m := make(map[string]string)
	for _, fld := range fe {
		m[fld.Field] = fld.Err
	}
	return m
}

func IsFieldErrors(err error) bool {
	var fe FieldErrors
	fmt.Println(errors.As(err, &fe))
	return errors.As(err, &fe)
}

func GetFieldErrors(err error) FieldErrors {
	var fe FieldErrors
	if !errors.As(err, &fe) {
		return nil
	}
	return fe
}
