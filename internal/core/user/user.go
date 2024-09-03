package user

import (
	"context"
	"fmt"
	"time"

	"github.com/aforamitdev/pdfsync/internal/core/event"
)

type UserRepository interface {
	WithInTran(ctx context.Context, fn func(s UserRepository) error) error
	Create(ctx context.Context, usr User) error
}

type Core struct {
	userRepo  UserRepository
	eventCore *event.Core
}

func NewCore(evnCore *event.Core, userRepo UserRepository) *Core {
	return &Core{userRepo: userRepo, eventCore: evnCore}
}

func (c *Core) Create(ctx context.Context, nu NewUser) (User, error) {

	now := time.Now()

	usr := User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: []byte(nu.Password),
		DateCreated:  now,
		DateUpdated:  now,
	}

	trans := func(s UserRepository) error {
		if err := s.Create(ctx, usr); err != nil {
			return fmt.Errorf("create %w", err)
		}
		return nil
	}

	if err := c.userRepo.WithInTran(ctx, trans); err != nil {
		return User{}, fmt.Errorf("tran: %w", err)
	}

	return usr, nil
}
