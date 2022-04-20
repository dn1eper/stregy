package stratexec

import (
	"context"
	"stregy/internal/domain/backtester"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) *repository {
	return &repository{db: client}
}

func (r *repository) CreateBacktester(ctx context.Context, bt backtester.Backtester, exchangeAccountID string) (*backtester.Backtester, error) {
	strategyIDParsed, _ := uuid.Parse(bt.Strategy.ID)
	exchangeAccountIDParsed, _ := uuid.Parse(exchangeAccountID)

	se := &StrategyExecution{
		StrategyID:        strategyIDParsed,
		ExchangeAccountID: exchangeAccountIDParsed,
		Timeframe:         bt.Timeframe,
		Symbol:            bt.Symbol,
		StartTime:         bt.StartDate,
		EndTime:           bt.EndDate,
		Status:            StrategyExecutionStatus(bt.Status),
	}
	result := r.db.Create(se)
	if result.Error != nil {
		return nil, result.Error
	}

	bt.ID = se.StrategyExecutionID.String()
	return &bt, nil
}
