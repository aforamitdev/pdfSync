package web

import (
	"context"
	"net/http"
)

type Encoder interface {
	Encode() ([]byte, error)
}

type Handler func(context.Context, http.ResponseWriter, *http.Request) (Encoder, error)

type Logger func(context.Context, string, ...any)

type App struct {
	log Logger
	mux *http.ServeMux
	mw  []Middleware
}
