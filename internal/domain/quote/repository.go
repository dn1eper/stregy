package quote

import (
	"time"
)

type Repository interface {
	Get(symbol string, startTime, endTime time.Time, limit, timeframeSec int) ([]Quote, error)
	Load(symbol, filePath, delimiter string, timeframeSec int) error
}
