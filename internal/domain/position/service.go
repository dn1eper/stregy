package position

import (
	"fmt"
	"stregy/internal/domain/order"

	"github.com/google/uuid"
)

type Service interface {
	Open(mainOrder *order.Order, takePrice *float64, stopPrice *float64, stratexecID string) (*Position, error)
	UpdateTakeProfit(posID string, orderID string, stratexecID string) error
	UpdateStopLoss(posID string, o *order.Order, stratexecID string) error
	Close(posID string, status PositionStatus) error
}

type service struct {
	repository   Repository
	orderService order.Service
}

func NewService(repository Repository, orderService order.Service) Service {
	return &service{
		repository:   repository,
		orderService: orderService,
	}
}

func (s *service) Open(mainOrder *order.Order, takePrice *float64, stopPrice *float64, stratexecID string) (*Position, error) {
	if *takePrice <= 0 || *stopPrice <= 0 {
		return nil, fmt.Errorf("takePrice and stopPrice must be greater than 0")
	}
	if mainOrder.Status == order.Partial {
		return nil, fmt.Errorf("partial positions not implemented")
	}
	if mainOrder.Status != order.Filled {
		return nil, fmt.Errorf("invalid order status")
	}

	var direction order.OrderDirection
	if mainOrder.Direction == order.Long {
		direction = order.Short
	} else {
		direction = order.Long
	}

	position := &Position{
		MainOrder: *mainOrder,
		Status:    Open,
	}
	if stopPrice != nil {
		o, err := s.orderService.Create(direction, mainOrder.Size, *stopPrice, order.StopMarket, mainOrder.SetupTime, stratexecID)
		if err != nil {
			return nil, err
		}
		position.StopOrder = o
	}
	if takePrice != nil {
		o, err := s.orderService.Create(direction, mainOrder.Size, *takePrice, order.StopMarket, mainOrder.SetupTime, stratexecID)
		if err != nil {
			return nil, err
		}
		position.TakeOrder = o
	}

	position, err := s.repository.Create(position)
	if err != nil {
		return nil, err
	}
	return position, nil
}

func (s *service) UpdateTakeProfit(posID string, takeOrderID string, stratexecID string) error {
	parsedTakeOrderID, _ := uuid.Parse(takeOrderID)
	return s.repository.Update(posID, map[string]interface{}{"take_order_id": parsedTakeOrderID})
}

func (s *service) UpdateStopLoss(posID string, o *order.Order, stratexecID string) error {
	return s.repository.Update(posID, map[string]interface{}{"stop_order_id": o})
}

func (s *service) Close(posID string, status PositionStatus) error {
	return s.repository.Update(posID, map[string]interface{}{"status": status})
}
