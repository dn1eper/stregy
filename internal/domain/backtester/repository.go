package backtester

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, backtester Backtester) (*Backtester, error)
}
