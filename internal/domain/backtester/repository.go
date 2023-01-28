package backtester

type Repository interface {
	Create(backtester Backtest) (*Backtest, error)
	GetBacktest(id string) (*Backtest, error)
}
