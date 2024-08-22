package user

import (
	"net/mail"
	"time"
)

type User struct {
	ID           int
	Name         string
	Email        mail.Address
	PasswordHash []byte
	Enabled      bool
	DateCreated  time.Time
	DateUpdated  time.Time
}

// informations needed to create a new user
type NewUser struct {
	Name  string
	Email mail.Address
}
