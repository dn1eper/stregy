package position

type PositionStatus int64

const (
	Created PositionStatus = iota
	Open
	TakeProfit
	StopLoss
	Cancelled
)

type Position struct {
	Status PositionStatus
}
