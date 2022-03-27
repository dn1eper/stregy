package order

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderStatus int64

const (
	Submitted OrderStatus = iota
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
	UUID           string
	Direction      OrderDirection
	Size           decimal.Decimal
	Price          decimal.Decimal
	Status         OrderStatus
	Type           OrderType
	ExecutionTime  time.Time
	ExecutionPrice decimal.Decimal
}
