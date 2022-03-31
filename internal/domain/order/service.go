package order

import (
	"context"

	"github.com/shopspring/decimal"
)

type Service interface {
	Create(ctx context.Context) (Order, error)
	Open(ctx context.Context, id string) error
	Fill(ctx context.Context, id string, size decimal.Decimal) (Order, error)
	Submit(ctx context.Context, id string) (Order, error)
}
