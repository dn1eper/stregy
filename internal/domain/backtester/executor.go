package backtester

import "context"

type Executor interface {
	Execute(ctx context.Context, b *Backtester) error
}
