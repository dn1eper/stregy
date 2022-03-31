package order

import (
	"stregy/internal/domain/quote"
	"time"

	"github.com/shopspring/decimal"
)

type OrderStatus int64

const (
	Draft OrderStatus = iota
	Submitted
	Accepted
	Partial
	Completed
	Cancelled
	Expired
	Margin
)

type OrderType int64

const (
	Limit OrderType = iota
	Market
	StopLimit
	StopMarket
	TrailingStop
)

type OrderDirection int64

const (
	Long OrderDirection = iota
	Short
)

type Order struct {
	ID             string
	Direction      OrderDirection
	Size           decimal.Decimal
	Price          decimal.Decimal
	Status         OrderStatus
	Type           OrderType
	ExecutionTime  time.Time
	ExecutionPrice decimal.Decimal
}

func (o Order) IsTouched(quote quote.Quote) bool {
	return o.Direction == Long && o.Price.LessThanOrEqual(quote.High) ||
		o.Direction == Short && o.Price.GreaterThanOrEqual(quote.Low)
}
