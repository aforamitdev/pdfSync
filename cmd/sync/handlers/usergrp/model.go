package usergrp

import (
	"fmt"
	"net/mail"

	"github.com/aforamitdev/pdfsync/internal/core/user"
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
		return user.NewUser{}, fmt.Errorf("parsing email %w", err)
	}
	user := user.NewUser{
		Name:  app.Name,
		Email: *addr,
	}

	return user, nil

}
