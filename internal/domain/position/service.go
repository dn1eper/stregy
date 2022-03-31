package position

import (
	"context"
	"stregy/internal/domain/order"

	"github.com/shopspring/decimal"
)

type Service interface {
	Create(ctx context.Context) (*Position, error)
	Open(ctx context.Context, position Position, size decimal.Decimal) (*Position, error)
	TakeProfit(ctx context.Context, position Position, size decimal.Decimal) (*Position, error)
	StopLoss(ctx context.Context, position Position, size decimal.Decimal) (*Position, error)
}

type service struct {
	repository   Repository
	orderService order.Service
}

func NewService(repository Repository, orderService order.Service) Service {
	return &service{repository: repository, orderService: orderService}
}

func (s *service) Create(ctx context.Context) (*Position, error) {
	panic("not implemented")
}

func (s *service) Open(ctx context.Context, position Position, size decimal.Decimal) (_ *Position, err error) {
	position.MainOrder, err = s.orderService.Fill(ctx, position.MainOrder.ID, size)
	if err != nil {
		return nil, err
	}

	position.TakeOrder, err = s.orderService.Submit(ctx, position.TakeOrder.ID)
	if err != nil {
		return nil, err
	}

	position.StopOrder, err = s.orderService.Submit(ctx, position.StopOrder.ID)
	if err != nil {
		return nil, err
	}

	if position.MainOrder.Status == order.Completed {
		position.Status = Open
	} else if position.MainOrder.Status == order.Partial {
		panic("TODO: partial position?")
	} else {
		panic("TODO: partial position?")
	}

	return &position, nil
}

func (s *service) TakeProfit(ctx context.Context, position Position, size decimal.Decimal) (*Position, error) {
	panic("not implemented")
}

func (s *service) StopLoss(ctx context.Context, position Position, size decimal.Decimal) (*Position, error) {
	panic("not implemented")
}
