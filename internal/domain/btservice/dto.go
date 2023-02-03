package btservice

import "time"

type BacktestDTO struct {
	StrategyName string
	SymbolName   string
	TimeframeSec int
	StartDate    time.Time
	EndDate      time.Time
}
