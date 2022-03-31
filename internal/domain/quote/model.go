package quote

import (
	"time"

	"github.com/shopspring/decimal"
)

type Quote struct {
	Time  time.Time
	Open  decimal.Decimal
	High  decimal.Decimal
	Low   decimal.Decimal
	Close decimal.Decimal
}
