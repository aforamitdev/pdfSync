package usergrp

import (
	"net/mail"

	"github.com/aforamitdev/pdfsync/internal/core/user"
	"github.com/aforamitdev/pdfsync/internal/sys/validate"
)

type AppNewUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
} //@name AppNewUser

func toCoreNewUser(app AppNewUser) (user.NewUser, error) {

	addr, err := mail.ParseAddress(app.Email)
	if err != nil {
		return user.NewUser{}, validate.NewFieldsError("email", err)
	}
	user := user.NewUser{
		Name:  app.Name,
		Email: *addr,
	}

	return user, nil

}

func (app AppNewUser) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil

}
