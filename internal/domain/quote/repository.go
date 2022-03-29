package quote

import (
	"context"
	"time"
)

type Repository interface {
	GetByIntervalPaginate(ctx context.Context, symbol string, start_time, end_time time.Time, offset, pageSize int) ([]Quote, error)
	Load(ctx context.Context, symbol string, filePath string, delimiter string) error
}
