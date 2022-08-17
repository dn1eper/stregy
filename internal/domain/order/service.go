package order

import (
	"fmt"
	"time"
)

type Service interface {
	Create(direction OrderDirection, size float64, price float64, orderType OrderType, setupTime time.Time, stratexecID string) (*Order, error)
	Get(orderID string) *Order
	ChangeSize(orderID string, size float64) error
	ChangePrice(orderID string, size float64) error
	Done(orderID string, status OrderStatus, doneTime time.Time, fillPrice float64) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(direction OrderDirection, size float64, price float64, orderType OrderType, setupTime time.Time, startexecID string) (*Order, error) {
	return s.repository.Create(&Order{
		Direction: direction,
		Size:      size,
		Price:     price,
		Type:      orderType,
		SetupTime: setupTime,
	}, startexecID)
}

func (s *service) Get(orderID string) *Order {
	return s.repository.Get(orderID)
}

func (s *service) ChangeSize(orderID string, size float64) error {
	if size <= 0 {
		return fmt.Errorf("size must be greater than 0")
	}
	return s.repository.Update(orderID, map[string]interface{}{"size": size})
}

func (s *service) ChangePrice(orderID string, price float64) error {
	if price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}
	return s.repository.Update(orderID, map[string]interface{}{"price": price})
}

func (s *service) Done(orderID string, status OrderStatus, doneTime time.Time, fillPrice float64) error {
	return s.repository.Update(orderID, map[string]interface{}{"status": status, "done_time": doneTime, "fill_price": fillPrice})
}
