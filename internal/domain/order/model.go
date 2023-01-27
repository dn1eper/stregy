package order

import (
	"stregy/internal/domain/quote"
	"time"
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
	Size           float64
	Price          float64
	Status         OrderStatus
	Type           OrderType
	ExecutionTime  time.Time
	ExecutionPrice float64
}

func (o Order) IsTouched(quote quote.Quote) bool {
	return o.Direction == Long && o.Price <= quote.High ||
		o.Direction == Short && o.Price <= quote.Low
}
