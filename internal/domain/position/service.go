package position

import (
	"stregy/internal/domain/order"
)

type Service interface {
	Create() (*Position, error)
	Open(position Position, size float64) (*Position, error)
	TakeProfit(position Position, size float64) (*Position, error)
	StopLoss(position Position, size float64) (*Position, error)
}

type service struct {
	repository   Repository
	orderService order.Service
}

func NewService(repository Repository, orderService order.Service) Service {
	return &service{repository: repository, orderService: orderService}
}

func (s *service) Create() (*Position, error) {
	panic("not implemented")
}

func (s *service) Open(position Position, size float64) (_ *Position, err error) {
	position.MainOrder, err = s.orderService.Fill(position.MainOrder.ID, size)
	if err != nil {
		return nil, err
	}

	position.TakeOrder, err = s.orderService.Submit(position.TakeOrder.ID)
	if err != nil {
		return nil, err
	}

	position.StopOrder, err = s.orderService.Submit(position.StopOrder.ID)
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

func (s *service) TakeProfit(position Position, size float64) (*Position, error) {
	panic("not implemented")
}

func (s *service) StopLoss(position Position, size float64) (*Position, error) {
	panic("not implemented")
}
