package order

import (
	"time"
)

type Order struct {
	ID             int64
	Diraction      OrderDiraction
	Size           float64
	Price          float64
	Status         OrderStatus
	Type           OrderType
	SubmissionTime time.Time
	FCTime         time.Time
	ExecutionPrice float64
	Position       *Position
}

//go:generate stringer -type=OrderStatus
type OrderStatus int64

const (
	Inactive OrderStatus = iota
	Submitted
	Accepted
	Partial
	Filled
	CancelledOrder
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

type OrderDiraction int64

const (
	Long OrderDiraction = iota
	Short
)

func (od OrderDiraction) Opposite() OrderDiraction {
	if od == Long {
		return Short
	}
	return Long
}

func (s OrderStatus) String() string {
	switch s {
	case Inactive:
		return "Inactive"
	case Submitted:
		return "Submitted"
	case Accepted:
		return "Accepted"
	case Partial:
		return "Partial"
	case Filled:
		return "Executed"
	case CancelledOrder:
		return "Cancelled"
	case Expired:
		return "Expired"
	case Margin:
		return "Margin"
	}

	return "uknown OrderStatus"
}

func (t OrderType) String() string {
	switch t {
	case Limit:
		return "Limit"
	case Market:
		return "Market"
	case StopLimit:
		return "StopLimit"
	case StopMarket:
		return "StopMarket"
	case TrailingStop:
		return "TrailingStop"
	}

	return "uknown OrderType"
}

func (t OrderDiraction) String() string {
	switch t {
	case Long:
		return "Long"
	case Short:
		return "Short"
	}

	return "uknown OrderDiraction"
}
