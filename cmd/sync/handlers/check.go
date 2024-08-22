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

type HealthResponse struct {
	Version   string `json:"version"`
	Status    string `json:"status"`
	Database  int    `json:"database"`
	DbVersion string `json:"db_version"`
} //@name HealthCheck

//	@Summary	health check
//	@Id			1
//	@version	1.0
//	@produce	application/json
//	@Success	200	{object}	HealthResponse	"health check response with, db connection, sqlite version,and app version "
//	@Router		/health [get]
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

	health := HealthResponse{

		Version:   c.build,
		Status:    "ok",
		Database:  c.db.Stats().OpenConnections,
		DbVersion: sqliteVersion,
	}
	return web.Response(ctx, w, health, http.StatusOK)

}
