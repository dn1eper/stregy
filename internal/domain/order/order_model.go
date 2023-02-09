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

type OrderStatus int64

const (
	InactiveOrder OrderStatus = iota
	SubmittedOrder
	AcceptedOrder
	PartialOrder
	FilledOrder
	CancelledOrder
	RejectedOrder
	ExpiredOrder
	MarginCallOrder
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

func (o *Order) Copy() *Order {
	return &Order{
		ID:             o.ID,
		Diraction:      o.Diraction,
		Size:           o.Size,
		Price:          o.Price,
		Status:         o.Status,
		Type:           o.Type,
		SubmissionTime: o.SubmissionTime,
		FCTime:         o.FCTime,
		ExecutionPrice: o.ExecutionPrice,
		Position:       o.Position}
}

func (od OrderDiraction) Opposite() OrderDiraction {
	if od == Long {
		return Short
	}
	return Long
}

func (s OrderStatus) String() string {
	switch s {
	case InactiveOrder:
		return "Inactive"
	case SubmittedOrder:
		return "Submitted"
	case AcceptedOrder:
		return "Accepted"
	case PartialOrder:
		return "Partial"
	case FilledOrder:
		return "Executed"
	case CancelledOrder:
		return "Cancelled"
	case ExpiredOrder:
		return "Expired"
	case MarginCallOrder:
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
