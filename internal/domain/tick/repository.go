package tick

import (
	"context"
	"time"
)

type Repository interface {
	GetByInterval(ctx context.Context, symbol string, startTime, endTime time.Time) ([]Tick, error)
	Load(symbol, filePath, delimiter string) error
}
