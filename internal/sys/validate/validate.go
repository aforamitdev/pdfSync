package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

// translator is a cache of locale and translation information.

func Check(val any) error {

	fmt.Println(val, "VALUE")
	if err := validate.Struct(val); err != nil {
		varrors, ok := err.(validator.ValidationErrors)

		if !ok {
			return err
		}

		var fields FieldErrors

		for _, verror := range varrors {
			field := FieldError{
				Field: verror.Field(),
				Err:   verror.Error(),
			}
			fields = append(fields, field)
		}
		return fields
	}
	return nil

}
