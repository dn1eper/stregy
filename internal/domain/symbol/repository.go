package symbol

import (
	"context"
)

type Repository interface {
	Exists(ctx context.Context, name string) bool
	Create(ctx context.Context, name string) (*Symbol, error)
	GetAll(ctx context.Context) ([]Symbol, error)
}
