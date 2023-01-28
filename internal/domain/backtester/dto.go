package backtester

import "time"

type BacktestDTO struct {
	StrategyName        string
	Timeframe           int
	Symbol              string
	StartDate           time.Time
	EndDate             time.Time
	HighOrderResolution bool
	BarsNeeded          int
	ATRperiod           int
}
