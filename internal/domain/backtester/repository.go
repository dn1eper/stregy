package backtester

import (
	"context"
)

type Repository interface {
	CreateBacktest(ctx context.Context, backtester Backtester) (*Backtester, error)
	GetBacktest(id string) (*Backtester, error)
}
