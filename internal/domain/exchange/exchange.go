package domain

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
)

type Broker interface {
	OnQuote(q *quote.Quote)
	OnOrder(o *order.Order)
	OnPosition(o *position.Position)
	OnExit()
}

type Exchange interface {
	RegisterOrder(o order.Order)
	CancelOrder(o order.Order)
	ClosePosition(p position.Position)
}

type exchange struct {
	broker *Broker
}

// TODO:
func (e exchange) RegisterOrder(o order.Order)       {}
func (e exchange) CancelOrder(o order.Order)         {}
func (e exchange) ClosePosition(p position.Position) {}

func NewExchange( /*TODO: params*/ ) Exchange {
	// TODO
	return exchange{}
}
