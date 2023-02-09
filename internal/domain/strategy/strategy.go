package strategy

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/quote"
	"time"
)

type Strategy interface {
	Name() string
	OnQuote(q quote.Quote, timeframe int)
	OnTick(price float64)
	OnOrder(o order.Order)

	QuoteTimeframesNeeded() []int
	TimeBeforeCallbacks() time.Duration // time to wait before first callback
}
