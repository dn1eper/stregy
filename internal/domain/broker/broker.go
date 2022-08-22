package broker

import (
	"stregy/internal/adapters/pgorm/exchange"
	"stregy/internal/domain/account"
	"stregy/internal/domain/dataseries"
	"stregy/internal/domain/order"
	"stregy/internal/domain/quote"
)

type Broker struct {
	ds  *dataseries.DataSeries
	acc *account.Account
	ex  *exchange.Exchange
}

func NewBroker(ds *dataseries.DataSeries, ac *account.Account, ex *exchange.Exchange) Broker {
	return Broker{ds, ac, ex}
}

func (b Broker) OnQuote(q *quote.Quote) {
	panic("not implemented")
}

func (b Broker) OnOrder(o *order.Order) {
	panic("not implemented")
}

func (b Broker) OnExit() {
	panic("not implemented")
}
