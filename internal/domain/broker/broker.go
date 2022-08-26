package broker

import (
	"stregy/internal/domain/dataseries"
	"stregy/internal/domain/exchange"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/strategy/core"
	"stregy/internal/domain/tradingac"
)

type Broker struct {
	dataSeries *dataseries.DataSeries
	account    *tradingac.Account
	exchange   *exchange.Exchange
	strategy   strategy.Implementation
}

func NewBroker() *Broker {
	return &Broker{}
}

func (b *Broker) Configure(
	ds *dataseries.DataSeries, acc *tradingac.Account,
	ex *exchange.Exchange, strat strategy.Implementation,
) {
	b.dataSeries = ds
	b.account = acc
	b.exchange = ex
	b.strategy = strat
}

func (b *Broker) FirstBars(q []quote.Quote) {
	ATRperiod := b.strategy.Config().ATRperiod
	// compute first ATR
	if ATRperiod != 0 {
		// quotes := make([]quote.Quote, 0, atrPeriod)
		// for len(quotes) != atrPeriod {
		// 	quotes = append(quotes, <-quoteChan)
		// }
		sum := 0.
		for _, q := range q[:ATRperiod] {
			sum += q.High - q.Low
		}
		core.ATR = sum / float64(ATRperiod)
	}
}

func (b *Broker) OnQuote(q *quote.Quote) {
}

func (b *Broker) OnOrder(o *order.Order) {
	panic("not implemented")
}

func (b *Broker) OnPosition(p *position.Position) {
	panic("not implemented")
}

func (b *Broker) OnExit() {
	panic("not implemented")
}

func (b *Broker) SendOrder(direction order.OrderDirection, size float64, price float64, orderType order.OrderType) (*order.Order, error) {
	panic("not implemented")
}

func (b *Broker) CancelOrder(orderID string) error {
	panic("not implemented")
}

func (b *Broker) ClosePosition(positionID string) error {
	panic("not implemented")
}
