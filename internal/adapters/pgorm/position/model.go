package position

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/order"
)

type Position struct {
	PositionID  uint64 `gorm:"primaryKey"`
	MainOrder   order.Order
	StopOrder   order.Order
	TakeOrder   order.Order
	MainOrderID uint64
	StopOrderID uint64
	TakeOrderID uint64
	Status      PositionStatus `gorm:"type:position_status;not null"`
}

type PositionStatus string

const (
	Created    PositionStatus = "CreatedPosition"
	Open       PositionStatus = "OpenPosition"
	TakeProfit PositionStatus = "TakeProfitPosition"
	StopLoss   PositionStatus = "StopLossPosition"
	Cancelled  PositionStatus = "CancelledPosition"
)

func (ps *PositionStatus) Scan(value interface{}) error {
	*ps = PositionStatus(value.([]byte))
	return nil
}

func (ps PositionStatus) Value() (driver.Value, error) {
	return string(ps), nil
}
