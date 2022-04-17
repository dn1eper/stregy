package backtester

import (
	"context"
)

type Repository interface {
	CreateBacktester(ctx context.Context, backtester Backtester, strategyID string, exchangeAccountID string) (*Backtester, error)
}
