package account

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
)

type Account struct {
	activeOrders    map[string]order.Order
	activePositions map[string]position.Position
}

func (a *Account) RegisterOrder(o *order.Order) {
	// create position if neccessary
	panic("not implemented")
}

func (a *Account) ActiveOrders() []order.Order {
	orders := make([]order.Order, 0, len(a.activeOrders))
	for _, order := range a.activeOrders {
		orders = append(orders, order)
	}
	return orders
}

func (a *Account) ActivePositions() []position.Position {
	positions := make([]position.Position, 0, len(a.activePositions))
	for _, pos := range a.activePositions {
		positions = append(positions, pos)
	}
	return positions
}
