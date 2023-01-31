package strat1

import (
	"stregy/internal/domain/bt"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/pkg/logging"
	"time"
)

var logger logging.Logger

type Strategy struct {
}

func NewStrategy() *Strategy {
	logger = logging.GetLogger()

	return &Strategy{}
}

func (s *Strategy) Name() string {
	return "strat1"
}

func (s *Strategy) OnOrder(order order.Order) {
}

func (s *Strategy) OnPosition(position position.Position) {
}

func (s *Strategy) OnQuote(quote quote.Quote, timeframe int) {
	bt.Printf("timeframe = %dm: %v", timeframe, quote)
}

func (s *Strategy) PrimaryTimeframeSec() int {
	return 1
}

func (s *Strategy) QuoteTimeframesNeeded() []int {
	return []int{5}
}

func (s *Strategy) TimeBeforeCallbacks() time.Duration {
	return time.Minute * 5 * 40
}

var _ strategy.Strategy = (*Strategy)(nil)
