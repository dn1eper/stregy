package exchange

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
)

type Broker interface {
	OnQuote(q *quote.Quote)
	OnOrder(o *order.Order)
	OnPosition(p *position.Position)
	OnExit()
	FirstBars(q []quote.Quote)
}

type Exchange interface {
	RegisterPosition(p *position.Position)
	CancelOrder(o *order.Order)
	ClosePosition(p *position.Position)
	Run() error
}
