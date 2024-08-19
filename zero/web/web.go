package web

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/gorilla/mux"
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

	handleFunc := wrapMiddleware(mw, handler)

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
