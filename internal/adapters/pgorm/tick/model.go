package tick

import (
	"time"
)

type Tick struct {
	Time  time.Time `gorm:"primaryKey;type:timestamp;index:time_idx_ticks"`
	Price float64   `gorm:"type:double precision"`
}
