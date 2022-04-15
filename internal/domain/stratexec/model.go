package stratexec

import "time"

type StrategyExecution struct {
	ID                string
	StrategyID        string
	ExchangeAccountID string
	Timeframe         int
	Symbol            string
	StartTime         time.Time
	EndTime           time.Time
}
