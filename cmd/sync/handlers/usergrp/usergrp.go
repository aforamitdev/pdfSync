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
// @Param		request	body		AppUser	true	"payload to ceate new app user"
// @Success	200		{object}	AppUser	" To create new user"
// @Router		/users [post]
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppUser
	if err := web.Decode(r, &app); err != nil {
		return web.NewRequestError(err, http.StatusBadRequest)

	}
	fmt.Println(app, "new")
	// nx, err := toNewAppUser(app)
	return nil
}
