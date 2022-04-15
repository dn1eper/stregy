package stratexec

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, se StrategyExecution) (*StrategyExecution, error)
}
