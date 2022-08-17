package position

import (
	"stregy/internal/domain/order"
)

type Position struct {
	PositionID string
	MainOrder  order.Order
	TakeOrder  *order.Order
	StopOrder  *order.Order
	Status     PositionStatus
}

type PositionStatus string

const (
	Open        PositionStatus = "OpenPosition"
	TakeProfit  PositionStatus = "TakeProfitPosition"
	StopLoss    PositionStatus = "StopLossPosition"
	MarketClose PositionStatus = "MarketClosePosition"
)
