package order

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/domain/order"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	StrategyExecution   stratexec.StrategyExecution
	StrategyExecutionID uuid.UUID
	Price               float64        `gorm:"type:double precision;not null"`
	Direction           OrderDirection `gorm:"type:order_direction"`
	Size                float64        `gorm:"type:double precision;not null"`
	Type                OrderType      `gorm:"type:order_type;not null"`
	SetupTime           time.Time      `gorm:"type:timestamp;not null"`
	DoneTime            time.Time      `gorm:"type:timestamp"`
	FillPrice           float64        `gorm:"type:double precision"`
	Status              OrderStatus    `gorm:"type:order_status;not null"`
}

type OrderType string

const (
	Limit         OrderType = "LimitOrder"
	Stop          OrderType = "StopOrder"
	Market        OrderType = "MarketOrder"
	StopLimit     OrderType = "StopLimitOrder"
	TrailingStop  OrderType = "TrailingStopOrder"
	CloseByLimit  OrderType = "CloseByLimitOrder"
	CloseByStop   OrderType = "CloseByStopOrder"
	CloseByMarket OrderType = "CloseByMarketOrder"
)

func (ot *OrderType) Scan(value interface{}) error {
	*ot = OrderType(value.(string))
	return nil
}

func (ot OrderType) Value() (driver.Value, error) {
	return string(ot), nil
}

type OrderStatus string

const (
	Submitted OrderStatus = "SubmittedOrder"
	Accepted  OrderStatus = "AcceptedOrder"
	Rejected  OrderStatus = "RejectedOrder"
	Partial   OrderStatus = "PartialOrder"
	Filled    OrderStatus = "FilledOrder"
	Cancelled OrderStatus = "CancelledOrder"
	Expired   OrderStatus = "ExpiredOrder"
	Margin    OrderStatus = "MarginOrder"
)

func (os *OrderStatus) Scan(value interface{}) error {
	*os = OrderStatus(value.(string))
	return nil
}

func (os OrderStatus) Value() (driver.Value, error) {
	return string(os), nil
}

type OrderDirection string

const (
	Short OrderDirection = "Short"
	Long  OrderDirection = "Long"
)

func (od *OrderDirection) Scan(value interface{}) error {
	*od = OrderDirection(value.(string))
	return nil
}

func (od OrderDirection) Value() (driver.Value, error) {
	return string(od), nil
}

func (o *Order) ToDomain() *order.Order {
	return &order.Order{
		OrderID:   o.OrderID.String(),
		Direction: order.OrderDirection(o.Direction),
		Price:     o.Price,
		Size:      o.Size,
		Type:      order.OrderType(o.Type),
		SetupTime: o.SetupTime,
		DoneTime:  o.DoneTime,
		FillPrice: o.FillPrice,
		Status:    order.OrderStatus(o.Status),
	}
}
