package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenSqlDb() (*sql.DB, error) {

	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		return nil, err
	}
	return db, nil

}
