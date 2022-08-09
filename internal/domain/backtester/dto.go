package backtester

import "time"

type BacktesterDTO struct {
	StrategyID          string
	Timeframe           int
	Symbol              string
	StartDate           time.Time
	EndDate             time.Time
	HighOrderResolution bool
	BarsNeeded          int
	ATRperiod           int
}
