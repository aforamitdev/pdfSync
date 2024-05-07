package handlers

import (
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type APIMuxConfig struct {
	shutdown chan os.Signal
	log      *zap.SugaredLogger
	DB       sqlx.DB
}

func APIMux(cfg APIMuxConfig) http.Handler {

}
