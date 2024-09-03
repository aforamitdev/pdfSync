package userdb

import (
	"net/mail"
	"time"

	"github.com/aforamitdev/pdfsync/internal/core/user"
)

type dbUser struct {
	Id          int64     `db:"id"`
	Name        string    `db:"user_name"`
	Email       string    `db:"user_email"`
	Password    string    `db:"password_hash"`
	DateCreate  time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

func toDBUser(usr user.User) dbUser {
	return dbUser{
		Name:        usr.Name,
		Email:       usr.Email.Address,
		Password:    string(usr.PasswordHash),
		DateCreate:  usr.DateCreated,
		DateUpdated: usr.DateUpdated,
	}

}

func toCoreUser(dbUsr dbUser) user.User {

	addr := mail.Address{
		Address: dbUsr.Email,
	}
	usr := user.User{
		ID:           int(dbUsr.Id),
		Name:         dbUsr.Name,
		Email:        addr,
		PasswordHash: []byte(dbUsr.Password),
	}
	return usr
}
