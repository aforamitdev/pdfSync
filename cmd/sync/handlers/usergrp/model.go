package usergrp

import (
	"net/mail"
	"time"

	"github.com/aforamitdev/pdfsync/internal/core/user"
)

type AppNewUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toNewAppUser(user AppNewUser) AppNewUser {


	adde,err:=mail.ParseAddress(user.Email)
	if err!=nil{
		return AppNewUser{}
	}


	return AppNewUser{
		Email:       user.Email.Address,
		Name:        user.Name,
		DateCreated: user.DateCreated.Format(time.RFC3339),
		DateUpdated: user.DateUpdated.Format(time.RFC3339)
	}
}
