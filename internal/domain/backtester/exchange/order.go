package exchange

import (
	"stregy/internal/domain/order"

	btree "github.com/ross-oreto/go-tree"
)

type Order struct {
	*order.Order
	abovePrice bool
}

func (o *Order) Comp(than btree.Val) int8 {
	if o.OrderID < than.(*Order).OrderID {
		return -1
	} else if o.OrderID > than.(*Order).OrderID {
		return 1
	}
	return 0
}
