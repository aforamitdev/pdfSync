package userdb

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aforamitdev/pdfsync/internal/core/user"
	"github.com/aforamitdev/pdfsync/internal/database"
	logger "github.com/aforamitdev/pdfsync/zero"
	"github.com/aforamitdev/pdfsync/zero/web"
)

type UserRepository struct {
	log    *logger.Logger
	db     *sql.DB
	inTran bool
}

func NewRepository(log *logger.Logger, db *sql.DB) *UserRepository {
	return &UserRepository{log: log, db: db}
}

func (s *UserRepository) Create(ctx context.Context, usr user.User) error {
	// fmt.Println(usr.Email, "USER RPO")
	q := `INSERT INTO users(user_name,email) VALUES(:user_name,:email)`
	q = database.QueryString(q, toDBUser(usr))
	fmt.Println(q, "QIERY")
	res, err := s.db.Exec(q)

	if err != nil {
		fmt.Print(err)
		return web.NewRequestError(fmt.Errorf("error inserting data"), http.StatusInternalServerError)
	}
	fmt.Println(res)
	return nil
}

func (s *UserRepository) WithInTran(ctx context.Context, fn func(s user.UserRepository) error) error {
	if s.inTran {
		return fn(s)
	}

	f := func(tx *sql.Tx, db *sql.DB) error {
		s := &UserRepository{log: s.log, db: db, inTran: true}

		return fn(s)
	}
	return database.WithInTran(ctx, s.log, s.db, f)

}
