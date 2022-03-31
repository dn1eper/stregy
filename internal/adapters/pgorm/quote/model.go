package quote

import (
	"stregy/internal/domain/quote"
	"time"

	"github.com/shopspring/decimal"
)

type Quote struct {
	Time  time.Time       `gorm:"primaryKey;type:timestamp;index:time_idx_quotes"`
	Open  decimal.Decimal `gorm:"type:decimal(20,8)"`
	High  decimal.Decimal `gorm:"type:decimal(20,8)"`
	Low   decimal.Decimal `gorm:"type:decimal(20,8)"`
	Close decimal.Decimal `gorm:"type:decimal(20,8)"`
}

func (q *Quote) ToDomain() quote.Quote {
	return quote.Quote{Time: q.Time, Open: q.Open, High: q.High, Low: q.Low, Close: q.Close}
}
