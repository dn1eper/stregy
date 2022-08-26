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

func (r *repository) CreateBacktest(ctx context.Context, bt backtester.Backtester) (*backtester.Backtester, error) {
	strategyIDParsed, _ := uuid.Parse(bt.Strategy.ID)

	se := &StrategyExecution{
		StrategyID:          strategyIDParsed,
		Timeframe:           bt.Timeframe,
		Symbol:              bt.Symbol,
		StartTime:           bt.StartDate,
		EndTime:             bt.EndDate,
		HighOrderResolution: bt.HighOrderResolution,
		Status:              StrategyExecutionStatus(bt.Status),
	}
	result := r.db.Create(se)
	if result.Error != nil {
		return nil, result.Error
	}

	bt.ID = se.StrategyExecutionID.String()
	return &bt, nil
}

func (r *repository) GetBacktest(id string) (*backtester.Backtester, error) {
	se := &StrategyExecution{}
	result := r.db.First(se, uuid.MustParse(id))
	if result.Error != nil {
		return nil, result.Error
	}

	bt := backtester.Backtester{
		ID:                  se.StrategyExecutionID.String(),
		Strategy:            *se.Strategy.ToDomain(),
		StartDate:           se.StartTime,
		EndDate:             se.EndTime,
		Symbol:              se.Symbol,
		Timeframe:           se.Timeframe,
		Status:              backtester.StrategyExecutionStatus(se.Status),
		HighOrderResolution: se.HighOrderResolution,
	}
	return &bt, nil
}
