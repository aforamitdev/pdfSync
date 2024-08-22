package handlers

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/aforamitdev/pdfsync/zero/web"
)

type check struct {
	build string
	db    *sql.DB
}

type HealthResponse struct {
	Version    string `json:"version"`
	Status     string `json:"status"`
	Database   int    `json:"database"`
	DbVersion  string `json:"db_version"`
	Host       string `json:"host"`
	GoMaxProcs string `json:"maxProcess"`
} //@name HealthCheck

// @Summary	health check
// @Id		1
// @version	1.0
// @produce	application/json
// @Success	200	{object}	HealthResponse	"health check response with, db connection, sqlite version,and app version "
// @Router	/health [get]
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

	host_name, err := os.Hostname()
	if err != nil {
		host_name = "N/A"
	}

	health := HealthResponse{

		Version:    c.build,
		Status:     "ok",
		Database:   c.db.Stats().OpenConnections,
		DbVersion:  sqliteVersion,
		Host:       host_name,
		GoMaxProcs: os.Getenv("GOMAXPROCS"),
	}
	return web.Response(ctx, w, health, http.StatusOK)

}
