package quote

import (
	"time"
)

type Quote struct {
	Time   time.Time `gorm:"primaryKey;type:timestamp"`
	Open   float64   `gorm:"double precision"`
	High   float64   `gorm:"double precision"`
	Low    float64   `gorm:"double precision"`
	Close  float64   `gorm:"double precision"`
	Volume int32     `gorm:"type:integer"`
}
