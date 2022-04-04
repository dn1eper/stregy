package stratexec

import (
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/adapters/pgorm/strategy"
	"time"

	"github.com/google/uuid"
)

type StrategyExecution struct {
	StrategyExecutionID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Strategy            strategy.Strategy
	StrategyID          uuid.UUID
	ExchangeAccount     exgaccount.ExchangeAccount
	ExchangeAccountID   uuid.UUID
	Timeframe           uint      `gorm:"type:int;not null;check:timeframe > 0"`
	Symbol              string    `gorm:"not null"`
	StartTime           time.Time `gorm:"type:timestamp;not null"`
	EndTime             time.Time `gorm:"type:timestamp"`
}
