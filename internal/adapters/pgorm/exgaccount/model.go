package exgaccount

import (
	"stregy/internal/adapters/pgorm/exchange"
	"stregy/internal/adapters/pgorm/user"

	"github.com/google/uuid"
)

type ExchangeAccount struct {
	ExchangeAccountID   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Exchange            exchange.Exchange
	User                user.User
	ExchangeID          uuid.UUID
	UserID              uuid.UUID
	ConnectionString    string `gorm:"type:text"`
	ExchangeAccountName string `gorm:"type:text;not null"`
}
