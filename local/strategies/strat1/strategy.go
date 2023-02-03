package strat1

import (
	"stregy/internal/domain/bt"
	"stregy/internal/domain/order"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/pkg/logging"
	"time"
)

var logger logging.Logger

var prevClose float64

type Strategy struct {
}

func NewStrategy() *Strategy {
	logger = logging.GetLogger()

	return &Strategy{}
}

func (s *Strategy) Name() string {
	return "strat1"
}

func (s *Strategy) OnOrder(o order.Order) {
	bt.PrintOrder(&o)
}

func (s *Strategy) OnQuote(q quote.Quote, timeframe int) {
	// bt.Printf("timeframe = %dm: %v", timeframe, quote)
	if prevClose == 0 {
		prevClose = q.Close
	} else if q.Close > prevClose {
		bt.SubmitContingentOrders(q.Close, 1, order.Long, order.Market, 100, 100)
	} else if q.Close < prevClose {
		bt.SubmitContingentOrders(q.Close, 1, order.Short, order.Market, 100, 100)
	}
	prevClose = q.Close
}

func (s *Strategy) PrimaryTimeframeSec() int {
	return 1
}

func (s *Strategy) QuoteTimeframesNeeded() []int {
	return []int{5}
}

func (s *Strategy) TimeBeforeCallbacks() time.Duration {
	return time.Minute * 0
}

var _ strategy.Strategy = (*Strategy)(nil)
