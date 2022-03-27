package strategy

import (
	"context"
)

type Repository interface {
	GetOne(ctx context.Context, uuid string) (*Strategy, error)
	Delete(ctx context.Context, uuid string) error
	Create(ctx context.Context, strat *Strategy) (*Strategy, error)
}
