package usergrp

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aforamitdev/pdfsync/internal/core/user"
	"github.com/aforamitdev/pdfsync/zero/web"
)

type Handlers struct {
	user *user.Core
}

func New(user *user.Core) *Handlers {
	return &Handlers{user: user}
}

// @Summary	create a new system user
// @Id			create-user
// @version	1.0
// @Produce	application/json
// @Param		request	body		AppNewUser	true	"payload to ceate new app user"
// @Success	200		{object}	AppNewUser	" To create new user"
// @Router		/users [post]
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewUser
	if err := web.Decode(r, &app); err != nil {
		return err

	}

	nc, err := toCoreNewUser(app)

	if err != nil {
		return err
	}

	fmt.Println(nc)
	return web.Response(ctx, w, app, http.StatusCreated)
}
