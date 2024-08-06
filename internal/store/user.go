package user

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return Store{db: db}
}

func (s Store) Create(ctx context.Context) {

}
