package usergrp

type AppUser struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

// func toAppUser(user AppUser) AppUser {

// 	adde, err := mail.ParseAddress(user.Email)
// 	if err != nil {
// 		return AppUser{}
// 	}

// 	return AppUser{
// 		Email:       user.Email,
// 		Name:        user.Name,
// 		DateCreated: user.DateCreated.Format(time.RFC3339),
// 		DateUpdated: user.DateUpdated.Format(time.RFC3339),
// 	}
// }
