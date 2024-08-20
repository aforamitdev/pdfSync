package user

import (
	"context"

	"github.com/aforamitdev/pdfsync/internal/core/event"
)

type Storer interface {
	Create(ctx context.Context, usr User)
}

type Core struct {
	storer    Storer
	eventCore *event.Core
}

func NewCore(evnCore *event.Core, storer Storer) *Core {
	return &Core{storer: storer, eventCore: evnCore}
}
