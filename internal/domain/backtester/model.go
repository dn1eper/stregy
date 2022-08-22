package backtester

import (
	"stregy/internal/domain/strategy"
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
	ID        string
	Strategy  strategy.Strategy
	StartDate time.Time
	EndDate   time.Time
	Symbol    string
	Timeframe int
	Status    StrategyExecutionStatus

	HighOrderResolution bool
}
