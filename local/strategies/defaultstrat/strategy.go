package defaultstrat

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"time"
)

type Strategy struct {
}

func NewStrategy() *Strategy {
	return &Strategy{}
}

func (s *Strategy) Name() string {
	return "defaultstrat"
}

func (s *Strategy) OnOrder(o order.Order) {
}

func (s *Strategy) OnQuote(q quote.Quote, timeframe int) {
}

func (s *Strategy) PrimaryTimeframeSec() int {
	return 60
}

func (s *Strategy) QuoteTimeframesNeeded() []int {
	return []int{}
}

func (s *Strategy) TimeBeforeCallbacks() time.Duration {
	return time.Minute * 5 * 14
}

var _ strategy.Strategy = (*Strategy)(nil)
