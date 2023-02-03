package stratexec

import (
	"database/sql/driver"
	"stregy/internal/adapters/pgorm/exgaccount"
	"stregy/internal/domain/bt"
	"stregy/internal/domain/symbol"
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
	StrategyExecutionId uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	StrategyName        string    `gorm:"type:varchar(100);not null"`
	ExchangeAccount     exgaccount.ExchangeAccount
	ExchangeAccountID   *uuid.UUID
	TimeframeSec        int                     `gorm:"type:int;not null;check:timeframe_sec > 0"`
	SymbolName          string                  `gorm:"not null"`
	StartTime           time.Time               `gorm:"type:timestamp;not null"`
	EndTime             time.Time               `gorm:"type:timestamp"`
	Status              StrategyExecutionStatus `gorm:"type:strategy_execution_status;not null"`
}

func (s *StrategyExecution) ToBacktest() *bt.Backtest {
	return &bt.Backtest{
		ID:           s.StrategyExecutionId.String(),
		StrategyName: s.StrategyName,
		StartTime:    s.StartTime,
		EndTime:      s.EndTime,
		Symbol:       symbol.Symbol{Name: s.SymbolName},
		TimeframeSec: s.TimeframeSec,
		Status:       bt.StrategyExecutionStatus(string(s.Status)),
	}
}
