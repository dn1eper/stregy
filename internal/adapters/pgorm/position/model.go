package position

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/order"
	"stregy/internal/domain/position"

	"github.com/google/uuid"
)

type Position struct {
	PositionID  uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	MainOrder   order.Order
	StopOrder   order.Order
	TakeOrder   order.Order
	MainOrderID uuid.UUID
	StopOrderID uuid.UUID
	TakeOrderID uuid.UUID
	Status      PositionStatus `gorm:"type:position_status;not null"`
}

type PositionStatus string

const (
	Open        PositionStatus = "OpenPosition"
	TakeProfit  PositionStatus = "TakeProfitPosition"
	StopLoss    PositionStatus = "StopLossPosition"
	MarketClose PositionStatus = "MarketClosePosition"
)

func (ps *PositionStatus) Scan(value interface{}) error {
	*ps = PositionStatus(value.(string))
	return nil
}

func (ps PositionStatus) Value() (driver.Value, error) {
	return string(ps), nil
}

func (p *Position) ToDomain() *position.Position {
	return &position.Position{
		PositionID: p.PositionID.String(),
		MainOrder:  *p.MainOrder.ToDomain(),
		StopOrder:  p.StopOrder.ToDomain(),
		TakeOrder:  p.TakeOrder.ToDomain(),
		Status:     position.PositionStatus(p.Status),
	}
}
