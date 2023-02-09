package bt

import (
	"stregy/internal/domain/order"
)

func (b *Backtester) GetBalance(stratexecID string) (float64, error) {
	return b.balance, nil
}

func (b *Backtester) updateBalance(o *order.Order) {
	if o == nil || o.Status != order.Filled || o.Position.MainOrder.ID == o.ID {
		return
	}

	p := o.Position
	if p.MainOrder.Diraction == order.Long {
		b.balance += o.ExecutionPrice - p.MainOrder.ExecutionPrice
	} else {
		b.balance += p.MainOrder.ExecutionPrice - o.ExecutionPrice
	}
}
