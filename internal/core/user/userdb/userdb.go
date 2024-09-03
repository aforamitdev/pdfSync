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
	fmt.Println(usr)
	q := `INSERT INTO system_users(id,user_name,user_email,password_hash,date_created,date_updated) VALUES(:id,:user_name,:user_email,:password_hash,:date_created,:date_updated)`
	q = database.QueryString(q, toDBUser(usr))
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
