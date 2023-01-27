package strategy

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
)

type Strategy interface {
	OnQuote(quote quote.Quote)
	OnOrder(order order.Order)
	OnPosition(position position.Position)
}
