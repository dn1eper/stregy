package tick

import (
	"time"

	"github.com/shopspring/decimal"
)

type Tick struct {
	Time  time.Time       `gorm:"primaryKey;type:timestamp;index:time_idx_ticks"`
	Price decimal.Decimal `gorm:"type:decimal(20,8)"`
}
