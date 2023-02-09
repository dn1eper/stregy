package bt

import "fmt"

func (b *Backtester) ChangeOrderPrice(id int64, price float64) error {
	if price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	o, ok := b.orders[id]
	if !ok {
		return fmt.Errorf("Order %d not found", id)
	}

	o.Price = price

	return nil
}
