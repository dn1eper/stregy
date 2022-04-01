package position

import (
	"stregy/internal/domain/order"

	"github.com/google/uuid"
)

type PositionStatus int

const (
	Created PositionStatus = iota
	Open
	TakeProfit
	StopLoss
	Cancelled
)

type Position struct {
	PositionID  uuid.UUID `gorm:"primaryKey;type:uuid"`
	MainOrder   order.Order
	TakeOrder   order.Order
	StopOrder   order.Order
	MainOrderID uuid.UUID
	StopOrderID uuid.UUID
	TakeOrderID uuid.UUID
	Status      PositionStatus
}
