package usergrp

import (
	"context"
	"net/http"

	"github.com/aforamitdev/pdfsync/zero/web"
)

type Handlers struct{}

func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewUser
	if err := web.Decode(r, &app); err != nil {
		return web.NewRequestError(err, http.StatusBadRequest)
	}

	nx, err := toNewAppUser(app)

}
