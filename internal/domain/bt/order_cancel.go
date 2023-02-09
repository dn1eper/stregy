package bt

import (
	"fmt"
	"stregy/internal/domain/order"
)

func (b *Backtester) CancelOrder(id int64) error {
	o, ok := b.orders[id]
	if !ok {
		return fmt.Errorf("order not found")
	}

	delete(b.orders, id)

	o.Status = order.CancelledOrder
	o.FCTime = b.curTime
	if o.ID != o.Position.MainOrder.ID {
		o.Position.RemoveCgtOrder(id)
	}

	b.logger.LogOrderStatusChange(o)
	b.strategy.OnOrder(*o)

	return nil
}

func (b *Backtester) cancelContingentOrders(o *order.Order) {
	for _, oCtg := range o.Position.CtgOrders {
		b.CancelOrder(oCtg.ID)
	}
}
