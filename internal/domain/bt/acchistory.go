package bt

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/symbol"
)

type AccountHistoryReport interface {
	CreateReport(orders []*order.Order, s symbol.Symbol, filePath string) error
}
