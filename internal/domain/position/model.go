package position

import (
	"stregy/internal/domain/order"
)

type PositionStatus int64

const (
	Created PositionStatus = iota
	Open
	TakeProfit
	StopLoss
	Cancelled
)

type Position struct {
	Id        string
	MainOrder order.Order
	TakeOrder order.Order
	StopOrder order.Order
	Status    PositionStatus
}
