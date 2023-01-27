package backtester

type Repository interface {
	Create(backtester Backtester) (*Backtester, error)
}
