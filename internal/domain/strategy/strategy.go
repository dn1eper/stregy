package strategy

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"time"
)

type Strategy interface {
	Name() string
	OnQuote(quote quote.Quote, timeframe int)
	OnOrder(order order.Order)
	OnPosition(position position.Position)

	QuoteTimeframesNeeded() []int
	TimeBeforeCallbacks() time.Duration // time to wait before first callback
}
