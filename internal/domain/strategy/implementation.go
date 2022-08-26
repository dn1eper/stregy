package strategy

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
)

type StrategyConfig struct {
	BarsNeeded int
	ATRperiod  int
}

type Implementation interface {
	OnQuote(quote quote.Quote)
	OnOrder(order order.Order)
	OnPosition(position position.Position)
	OnExit()
	Config() StrategyConfig
}

type Broker interface {
	SendOrder(direction order.OrderDirection, size float64, price float64, orderType order.OrderType) (*order.Order, error)
	CancelOrder(orderID string) error
	ClosePosition(positionID string) error
}
