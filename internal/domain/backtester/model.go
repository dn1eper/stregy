package backtester

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"time"
)

type StrategyExecutionStatus string

const (
	Created  StrategyExecutionStatus = "Created"
	Running  StrategyExecutionStatus = "Running"
	Finished StrategyExecutionStatus = "Finished"
	Crashed  StrategyExecutionStatus = "Crashed"
)

type Backtester struct {
	ID           string
	StrategyName string
	StartDate    time.Time
	EndDate      time.Time
	Symbol       string
	Timeframe    int
	Status       StrategyExecutionStatus

	HighOrderResolution bool
	BarsNeeded          int
	Quotes              []quote.Quote

	activeOrders    []order.Order
	activePositions []position.Position

	ATRperiod int
	ATR       float64
}
