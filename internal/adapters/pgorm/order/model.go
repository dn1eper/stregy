package order

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/stratexec"
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID             uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	StrategyExecution   stratexec.StrategyExecution
	StrategyExecutionID uuid.UUID
	Price               float64        `gorm:"double precision;not null"`
	Direction           OrderDirection `gorm:"type:order_direction"`
	Size                float64        `gorm:"double precision;not null"`
	Type                OrderType      `gorm:"type:order_type;not null"`
	ExecutionTime       time.Time      `gorm:"type:timestamp"`
	ExecutionPrice      float64        `gorm:"double precision"`
	Status              OrderStatus    `gorm:"type:order_status;not null"`
}

type OrderType string

const (
	Limit        OrderType = "LimitOrder"
	Market       OrderType = "MarketOrder"
	StopLimit    OrderType = "StopLimitOrder"
	StopMarket   OrderType = "StopMarketOrder"
	TrailingStop OrderType = "TrailingStopOrder"
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
	Completed OrderStatus = "CompletedOrder"
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
