package exchange

import "github.com/google/uuid"

type Exchange struct {
	ExchangeID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string    `gorm:"type:text"`
}
