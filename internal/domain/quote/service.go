package quote

import (
	"context"
	"time"
)

type Service interface {
	GetByInterval(ctx context.Context, symbol string, start, end time.Time) ([]*Quote, error)
}
