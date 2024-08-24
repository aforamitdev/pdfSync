package validate

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// validate holds the settings and caches for validating request struct values.
var validate *validator.Validate

// translator is a cache of locale and translation information.

func init() {

	fmt.Println("INIT")
	validate = validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func Check(val any) error {

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
