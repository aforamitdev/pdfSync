package user

import "context"

type Storer interface {
	Create(ctx context.Context, usr User)
}
