package handlers

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/aforamitdev/pdfsync/zero/web"
)

type check struct {
	build string
	db    *sql.DB
}

// @Summary		Greeter service
// @Id			1
// @version		1.0
// @produce		application/json
// @Param		name	query		string	true	"Input name"
// @Success		200
// @Router		/health [get]
func (c *check) health(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var sqliteVersion string

	err := c.db.QueryRow("SELECT sqlite_version();").Scan(&sqliteVersion)

	if err != nil {
		if err == sql.ErrNoRows {
			web.NewRequestError(err, http.StatusNoContent)
			return err
		}

		web.ResponseError(ctx, w, err)
		return err
	}

	health := struct {
		Version   string `json:"version"`
		Status    string `json:"status"`
		Database  int    `json:"database"`
		DbVersion string `json:"db_version"`
	}{
		Version:   c.build,
		Status:    "ok",
		Database:  c.db.Stats().OpenConnections,
		DbVersion: sqliteVersion,
	}
	return web.Response(ctx, w, health, http.StatusOK)

}

// @Summary		test service
// @Id			2
// @version		1.0
// @produce		application/json
// @Param		name	query		string	true	"Input name"
// @Success		200
// @Router		/healthdv [get]
func (c *check) healdbth(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var sqliteVersion string

	err := c.db.QueryRow("SELECT sqlite_version();").Scan(&sqliteVersion)

	if err != nil {
		if err == sql.ErrNoRows {
			web.NewRequestError(err, http.StatusNoContent)
			return err
		}

		web.ResponseError(ctx, w, err)
		return err
	}

	health := struct {
		Version   string `json:"version"`
		Status    string `json:"status"`
		Database  int    `json:"database"`
		DbVersion string `json:"db_version"`
	}{
		Version:   c.build,
		Status:    "ok",
		Database:  c.db.Stats().OpenConnections,
		DbVersion: sqliteVersion,
	}
	return web.Response(ctx, w, health, http.StatusOK)

}
