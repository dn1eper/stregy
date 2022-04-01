package exchange

import "github.com/google/uuid"

type Exchange struct {
	ExchangeID uuid.UUID `gorm:"primaryKey;type:uuid"`
	Name       string    `gorm:"type:text"`
}
