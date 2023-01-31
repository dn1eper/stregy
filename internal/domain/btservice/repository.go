package btservice

import "stregy/internal/domain/bt"

type Repository interface {
	Create(backtester bt.Backtest) (*bt.Backtest, error)
	GetBacktest(id string) (*bt.Backtest, error)
}
