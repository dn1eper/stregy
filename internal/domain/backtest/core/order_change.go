package core

import "fmt"

func (b *Backtest) ChangeOrderPrice(id int64, price float64) error {
	if price <= 0 {
		return fmt.Errorf("price must be greater than 0")
	}

	o, ok := b.orders[id]
	if !ok {
		return &OrderNotFoundError{id}
	}

	o.Price = price

	return nil
}
