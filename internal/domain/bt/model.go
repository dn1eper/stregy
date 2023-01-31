package bt

import (
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
	TimeframeSec int
	Status       StrategyExecutionStatus
}
