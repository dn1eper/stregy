package quote

import (
	"stregy/internal/domain/quote"
	"time"
)

type Quote struct {
	Time  time.Time `gorm:"primaryKey;type:timestamp;index:time_idx_quotes"`
	Open  float32   `gorm:"type:real"`
	High  float32   `gorm:"type:real"`
	Low   float32   `gorm:"type:real"`
	Close float32   `gorm:"type:real"`
}

func (q *Quote) ToDomain() quote.Quote {
	return quote.Quote{Time: q.Time, Open: q.Open, High: q.High, Low: q.Low, Close: q.Close}
}
