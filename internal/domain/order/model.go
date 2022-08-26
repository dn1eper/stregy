package order

import (
	"time"
)

type Order struct {
	OrderID    string
	Direction  OrderDirection
	Size       float64
	Price      float64
	Status     OrderStatus
	Type       OrderType
	SetupTime  time.Time
	DoneTime   time.Time
	FillPrice  float64
	PositionID string
}

type OrderStatus string

const (
	Draft     OrderStatus = "Draft"
	Submitted OrderStatus = "SubmittedOrder"
	Accepted  OrderStatus = "AcceptedOrder"
	Rejected  OrderStatus = "RejectedOrder"
	Partial   OrderStatus = "PartialOrder"
	Filled    OrderStatus = "FilledOrder"
	Cancelled OrderStatus = "CancelledOrder"
	Expired   OrderStatus = "ExpiredOrder"
	Margin    OrderStatus = "MarginOrder"
)

type OrderType string

const (
	Limit         OrderType = "LimitOrder"
	Market        OrderType = "MarketOrder"
	Stop          OrderType = "StopOrder"
	StopLimit     OrderType = "StopLimitOrder"
	TrailingStop  OrderType = "TrailingStopOrder"
	CloseByLimit  OrderType = "CloseByLimitOrder"
	CloseByStop   OrderType = "CloseByStopOrder"
	CloseByMarket OrderType = "CloseByMarketOrder"
)

type OrderDirection string

const (
	Long  OrderDirection = "Long"
	Short OrderDirection = "Short"
)

func OppositeDirection(direction OrderDirection) OrderDirection {
	if direction == Long {
		return Short
	}
	return Long
}
