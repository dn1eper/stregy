package order

type Position struct {
	ID        int64
	MainOrder *Order
	TakeOrder *Order
	StopOrder *Order
	Status    PositionStatus
}

type PositionStatus int

const (
	CreatedPosition PositionStatus = iota
	PartialPosition
	OpenPosition
	TakeProfit
	StopLoss
	CancelledPosition
)
