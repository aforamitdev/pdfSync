package handlers

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/aforamitdev/pdfsync/cmd/sync/handlers/usergrp"
	"github.com/aforamitdev/pdfsync/internal/core/event"
	"github.com/aforamitdev/pdfsync/internal/core/user"
	"github.com/aforamitdev/pdfsync/internal/core/user/userdb"
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

	usrCore := user.NewCore(envCore, userdb.NewRepository(log, db))

	ugh := usergrp.New(usrCore)

	app := web.NewApp(shutdown, mid.Error(log))
	app.RegisterSwagger()

	check := check{build: build, db: db}

	app.Handle(http.MethodGet, "/v1/health", check.health)

	// user handlers
	app.Handle(http.MethodPost, "/v1/users", ugh.Create)

	return app
}
