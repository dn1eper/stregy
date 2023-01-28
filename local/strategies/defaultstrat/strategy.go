package defaultstrat

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
)

type Strategy struct {
}

func (s *Strategy) OnOrder(order order.Order) {
}

func (s *Strategy) OnPosition(position position.Position) {
}

func (s *Strategy) OnQuote(quote quote.Quote) {
}

var _ strategy.Strategy = (*Strategy)(nil)
