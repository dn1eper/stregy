package stratexec

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/adapters/pgorm/strategy"
	"stregy/internal/domain/stratexec"
	"time"

	"github.com/google/uuid"
)

type StrategyExecutionStatus stratexec.StrategyExecutionStatus

func (se *StrategyExecutionStatus) Scan(value interface{}) error {
	*se = StrategyExecutionStatus(value.(string))
	return nil
}

func (se StrategyExecutionStatus) Value() (driver.Value, error) {
	return string(se), nil
}

type StrategyExecution struct {
	StrategyExecutionID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Strategy            strategy.Strategy
	StrategyID          uuid.UUID
	ExchangeAccount     exgaccount.ExchangeAccount
	ExchangeAccountID   uuid.UUID
	Timeframe           int                     `gorm:"type:int;not null;check:timeframe > 0"`
	Symbol              string                  `gorm:"not null"`
	StartTime           time.Time               `gorm:"type:timestamp;not null"`
	EndTime             time.Time               `gorm:"type:timestamp"`
	Status              StrategyExecutionStatus `gorm:"type:strategy_execution_status;not null"`
}

func (se StrategyExecution) ToDomain() *stratexec.StrategyExecution {
	seDomain := stratexec.StrategyExecution{
		ID:                se.StrategyExecutionID.String(),
		StrategyID:        se.StrategyID.String(),
		ExchangeAccountID: se.ExchangeAccountID.String(),
		Timeframe:         se.Timeframe,
		Symbol:            se.Symbol,
		StartDate:         se.StartTime,
		EndDate:           se.EndTime,
		Status:            stratexec.StrategyExecutionStatus(se.Status),
	}
	return &seDomain
}
