package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type validator interface {
	Validate() error
}

func Decode(r *http.Request, val any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(val); err != nil {
		return err
	}
	fmt.Println("", "s")
	if v, ok := val.(validator); ok {
		fmt.Println(v.Validate(), "CAST SUCCES")
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
