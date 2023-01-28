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

type Backtest struct {
	Id           string
	StrategyName string
	StartTime    time.Time
	EndTime      time.Time
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
