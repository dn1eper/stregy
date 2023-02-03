package bt

import (
	"stregy/internal/domain/symbol"
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
	ID                    string
	StrategyName          string
	StartTime             time.Time
	EndTime               time.Time
	Symbol                symbol.Symbol
	TimeframeSec          int
	Status                StrategyExecutionStatus
	AccountHistoryService AccountHistoryReport
}
