package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	logger "github.com/aforamitdev/pdfsync/zero"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func OpenSqlDb() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		return nil, err
	}
	return db, nil

}

func WithInTran(ctx context.Context, log *logger.Logger, db *sql.DB, fn func(*sql.Tx, *sql.DB) error) error {

	log.Info(ctx, "begain trans")

	tx, err := db.Begin()

	if err != nil {
		return fmt.Errorf("begin tran: %w ", err)
	}

	defer func() {
		if err := tx.Rollback(); err != nil {
			if errors.Is(err, sql.ErrTxDone) {
				return
			}
			log.Errorw(ctx, "unable to run transections", "ERROR", err)
		}
		log.Info(ctx, "rollback tran")
	}()

	if err := fn(tx, db); err != nil {
		return fmt.Errorf("exec tran:%w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tran: %w", err)
	}
	log.Info(ctx, "commited translections")

	return nil
}

func QueryString(query string, args any) string {
	query, params, err := sqlx.Named(query, args)
	if err != nil {
		return err.Error()
	}

	for _, param := range params {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("'%s'", v)
		case []byte:
			value = fmt.Sprintf("'%s'", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}
		query = strings.Replace(query, "?", value, 1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.Trim(query, " ")
}
