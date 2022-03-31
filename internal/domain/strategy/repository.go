package strategy

import (
	"context"
)

type Repository interface {
	GetOne(ctx context.Context, id string) (*Strategy, error)
	Delete(ctx context.Context, id string) error
	Create(ctx context.Context, strategy Strategy) (*Strategy, error)
}
