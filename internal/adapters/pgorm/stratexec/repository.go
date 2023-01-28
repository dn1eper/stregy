package stratexec

import (
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

func (r *repository) Create(bt backtester.Backtest) (*backtester.Backtest, error) {
	se := &StrategyExecution{
		StrategyName:        bt.StrategyName,
		Timeframe:           bt.Timeframe,
		Symbol:              bt.Symbol,
		StartTime:           bt.StartTime,
		EndTime:             bt.EndTime,
		HighOrderResolution: bt.HighOrderResolution,
		Status:              StrategyExecutionStatus(bt.Status),
	}
	result := r.db.Create(se)
	if result.Error != nil {
		return nil, result.Error
	}

	bt.Id = se.StrategyExecutionId.String()
	return &bt, nil
}

func (r *repository) Get(id string) (*StrategyExecution, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	strategyExecution := &StrategyExecution{StrategyExecutionId: parsed}
	result := r.db.First(strategyExecution)

	return strategyExecution, result.Error
}

func (r *repository) GetBacktest(id string) (*backtester.Backtest, error) {
	strategyExecution, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	return strategyExecution.ToBacktest(), err
}
