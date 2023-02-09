package btservice

import "stregy/internal/domain/bt"

type Repository interface {
	Create(backtester bt.Backtester) (*bt.Backtester, error)
	GetBacktest(id string) (*bt.Backtester, error)
}
