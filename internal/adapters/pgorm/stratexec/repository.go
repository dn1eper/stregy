package stratexec

import (
	"context"
	"stregy/internal/domain/stratexec"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) stratexec.Repository {
	return &repository{db: client}
}

func (r *repository) Create(ctx context.Context, se stratexec.StrategyExecution) (*stratexec.StrategyExecution, error) {
	strategyID, _ := uuid.Parse(se.StrategyID)
	exchangeAccountID, _ := uuid.Parse(se.ExchangeAccountID)

	db_se := &StrategyExecution{
		StrategyID:        strategyID,
		ExchangeAccountID: exchangeAccountID,
		Timeframe:         se.Timeframe,
		Symbol:            se.Symbol,
		StartTime:         se.StartDate,
		EndTime:           se.EndDate,
		Status:            StrategyExecutionStatus(se.Status),
	}
	result := r.db.Create(db_se)
	return db_se.ToDomain(), result.Error
}
