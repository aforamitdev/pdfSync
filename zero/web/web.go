package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type Encoder interface {
	Encode() ([]byte, error)
}

type Handler func(context.Context, http.ResponseWriter, *http.Request) error

type Logger func(context.Context, string, ...any)

type App struct {
	mux      *mux.Router
	mw       []Middleware
	shutdown chan os.Signal
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {

	mux := mux.NewRouter()

	return &App{mux: mux, shutdown: shutdown, mw: mw}

}

func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {

	handleFunc := wrapMiddleware(a.mw, handler)
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		if err := handleFunc(ctx, w, r); err != nil {
			fmt.Println("error", err)
		}
	}

	a.mux.HandleFunc(path, h).Methods(method)
}

func (a *App) handleShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) RegisterSwagger() {

	a.mux.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)
}
