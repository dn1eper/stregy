package backtester

import (
	"stregy/internal/domain/strategy"
	"time"
)

type Backtester struct {
	ID        string
	Strategy  strategy.Strategy
	StartDate time.Time
	EndDate   time.Time
	Symbol    string
	Timeframe int
}
