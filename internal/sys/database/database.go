package database

import (
	"errors"
	"net/url"

	"github.com/jmoiron/sqlx"
)

var (
	ErrNotFound            = errors.New("not found")
	ErrInvalidID           = errors.New("ID is not in it proper form")
	ErrAuthenticateFailure = errors.New("authenticate failed")
	ErrForbidden           = errors.New("attempted action is not allowed")
)

type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}

func Open(cfg Config) (*sqlx.DB, error) {
	sslMode := "require"

	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")

	u := url.URL{Scheme: "postgres", User: url.UserPassword(cfg.User, cfg.Password), Host: cfg.Host, Path: cfg.Name, RawQuery: q.Encode()}
	db, err := sqlx.Open("postgres", u.String())

	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxIdleConns(cfg.MaxOpenConns)
	return db, nil
}
