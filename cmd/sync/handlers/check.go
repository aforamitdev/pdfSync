package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/aforamitdev/pdfsync/zero/web"
)

type check struct {
	build string
	db    *sql.DB
}

func (c *check) health(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var sqliteVersion int

	row, err := c.db.Query("SELECT sqlite_version();")
	if err != nil {
		web.ResponseError(ctx, w, err)
	}
	fmt.Println(row.Columns())
	err = row.Scan(sqliteVersion)
	if err != nil {
		web.ResponseError(ctx, w, err)
	}
	health := struct {
		Version   string `json:"version"`
		Status    string `json:"status"`
		Database  int    `json:"database"`
		DbVersion int    `json:"dbVersion"`
	}{
		Version:   c.build,
		Status:    "ok",
		Database:  c.db.Stats().OpenConnections,
		DbVersion: sqliteVersion,
	}
	return web.Response(ctx, w, health, http.StatusOK)

}
