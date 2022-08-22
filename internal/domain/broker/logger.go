package broker

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
)

type Logger interface {
	LogOrder(o *order.Order) error
	LogPosition(p *position.Position) error
}

type logger struct {
	orderService    *order.Service
	positionService *position.Service
}

func (l *logger) LogPosition(p *position.Position) error {
	return nil
}

func (l *logger) LogOrder(o *order.Order) error {
	return nil
}
