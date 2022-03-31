package quote

import (
	"context"
	"time"
)

type Repository interface {
	GetByInterval(ctx context.Context, symbol string, start_time, end_time time.Time, offset, pageSize int) ([]Quote, error)
}
