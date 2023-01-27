package stratexec

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/exgaccount"
	"time"

	"github.com/google/uuid"
)

type StrategyExecutionStatus string

func (se *StrategyExecutionStatus) Scan(value interface{}) error {
	*se = StrategyExecutionStatus(value.(string))
	return nil
}

func (se StrategyExecutionStatus) Value() (driver.Value, error) {
	return string(se), nil
}

type StrategyExecution struct {
	StrategyExecutionID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ExchangeAccount     exgaccount.ExchangeAccount
	ExchangeAccountID   *uuid.UUID
	Timeframe           int                     `gorm:"type:int;not null;check:timeframe > 0"`
	Symbol              string                  `gorm:"not null"`
	StartTime           time.Time               `gorm:"type:timestamp;not null"`
	EndTime             time.Time               `gorm:"type:timestamp"`
	HighOrderResolution bool                    `gorm:"type:bool"`
	Status              StrategyExecutionStatus `gorm:"type:strategy_execution_status;not null"`
}
