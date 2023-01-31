package btservice

import "time"

type BacktestDTO struct {
	StrategyName string
	Symbol       string
	TimeframeSec int
	StartDate    time.Time
	EndDate      time.Time
}
