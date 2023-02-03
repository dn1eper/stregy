package bt

import (
	"strconv"
	"stregy/internal/domain/order"
)

func PrintOrder(o *order.Order) {
	price := strconv.FormatFloat(o.Price, 'f', loggingConfig.PricePrecision, 64)
	Printf("Order #%d %s: %s %s", o.ID, o.Status.String(), price, o.Diraction.String())
}

func PrintOrderStatus(o *order.Order) {
	Printf("Order #%d: %s", o.ID, o.Status.String())
}
