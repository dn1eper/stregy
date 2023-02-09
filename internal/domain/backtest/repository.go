package backtest

import "stregy/internal/domain/backtest/core"

type Repository interface {
	Create(backtest core.Backtest) (*core.Backtest, error)
	GetBacktest(id string) (*core.Backtest, error)
}
