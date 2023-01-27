package quote

import (
	"time"
)

type Repository interface {
	GetByInterval(symbol string, start_time, end_time time.Time) ([]Quote, error)
	Load(symbol, filePath, delimiter string, timeframe string) error
}
