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

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Create(ctx context.Context) (Order, error) {
	panic("not implemented")
}

func (s *service) Open(ctx context.Context, id string) error {
	panic("not implemented")
}

func (s *service) Fill(ctx context.Context, id string, size decimal.Decimal) (Order, error) {
	panic("not implemented")
}

func (s *service) Submit(ctx context.Context, id string) (Order, error) {
	panic("not implemented")
}
