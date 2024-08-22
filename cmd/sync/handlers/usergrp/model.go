package usergrp

type AppUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
} //@name AppUser
