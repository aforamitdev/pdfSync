package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/aforamitdev/pdfsync/internal/core/event"
	"github.com/aforamitdev/pdfsync/internal/core/user"
	"github.com/aforamitdev/pdfsync/internal/mid"
	logger "github.com/aforamitdev/pdfsync/zero"
	"github.com/aforamitdev/pdfsync/zero/web"
)

type APIMuxConfig struct {
	shutdown chan os.Signal
	log      logger.Logger
	DB       *sql.DB
}

func APIMux(build string, shutdown chan os.Signal, log *logger.Logger, db *sql.DB) http.Handler {

	envCore := event.NewCore(log)

	usrCore := user.NewCore(envCore)

	app := web.NewApp(shutdown, mid.Error(log))
	check := check{build: build, db: db}
	app.Handle(http.MethodGet, "/v1/health", check.health)

	return app
}
