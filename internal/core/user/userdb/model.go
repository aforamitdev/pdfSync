package userdb

import (
	"net/mail"

	"github.com/aforamitdev/pdfsync/internal/core/user"
)

type dbUser struct {
	Id    int64  `db:"id"`
	Name  string `db:"user_name"`
	Email string `db:"email"`
}

func toDBUser(usr user.User) dbUser {
	return dbUser{
		Id:    1,
		Name:  usr.Name,
		Email: usr.Email.Address,
	}

}

func toCoreUser(dbUsr dbUser) user.User {

	addr := mail.Address{
		Address: dbUsr.Email,
	}
	usr := user.User{
		ID:    int(dbUsr.Id),
		Name:  dbUsr.Name,
		Email: addr,
	}
	return usr
}
