package strategy

import (
	"context"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
)

type Implementation interface {
	OnQuote(ctx context.Context, quote quote.Quote)
	OnOrder(ctx context.Context, order order.Order)
	OnPosition(ctx context.Context, position position.Position)
}
