package tick

import (
	"time"
)

type Repository interface {
	GetByInterval(symbol string, startTime, endTime time.Time) ([]Tick, error)
	Load(symbol, filePath, delimiter string) error
}
