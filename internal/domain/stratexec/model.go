package stratexec

import "time"

type StrategyExecutionStatus string

const (
	Created  StrategyExecutionStatus = "Created"
	Running  StrategyExecutionStatus = "Running"
	Finished StrategyExecutionStatus = "Finished"
	Crashed  StrategyExecutionStatus = "Crashed"
)

type StrategyExecution struct {
	ID                string
	StrategyID        string
	ExchangeAccountID string
	Timeframe         int
	Symbol            string
	StartDate         time.Time
	EndDate           time.Time
	Status            StrategyExecutionStatus
}
