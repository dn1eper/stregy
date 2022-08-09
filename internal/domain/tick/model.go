package tick

import (
	"time"

	"github.com/shopspring/decimal"
)

type Tick struct {
	Time  time.Time
	Price decimal.Decimal
}
