package userdb

import (
	"context"
	"database/sql"

	"github.com/aforamitdev/pdfsync/internal/core/user"
	logger "github.com/aforamitdev/pdfsync/zero"
)

type Store struct {
	log *logger.Logger
	db  *sql.DB
}

func NewStore(log *logger.Logger, db *sql.DB) *Store {
	return &Store{log: log, db: db}
}

func (s *Store) Create(ctx context.Context, usr user.User) error {
	return nil
}
