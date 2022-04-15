package stratexec

import (
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/adapters/pgorm/strategy"
	"stregy/internal/domain/stratexec"
	"time"

	"github.com/google/uuid"
)

type StrategyExecution struct {
	StrategyExecutionID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Strategy            strategy.Strategy
	StrategyID          uuid.UUID
	ExchangeAccount     exgaccount.ExchangeAccount
	ExchangeAccountID   uuid.UUID
	Timeframe           int       `gorm:"type:int;not null;check:timeframe > 0"`
	Symbol              string    `gorm:"not null"`
	StartTime           time.Time `gorm:"type:timestamp;not null"`
	EndTime             time.Time `gorm:"type:timestamp"`
}

func (se StrategyExecution) ToDomain() *stratexec.StrategyExecution {
	seDomain := stratexec.StrategyExecution{
		ID:                se.StrategyExecutionID.String(),
		StrategyID:        se.StrategyID.String(),
		ExchangeAccountID: se.ExchangeAccountID.String(),
		Timeframe:         se.Timeframe,
		Symbol:            se.Symbol,
		StartTime:         se.StartTime,
		EndTime:           se.EndTime,
	}
	return &seDomain
}
