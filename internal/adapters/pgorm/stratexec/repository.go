package stratexec

import (
	"stregy/internal/domain/bt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) *repository {
	return &repository{db: client}
}

func (r *repository) Create(backtest bt.Backtest) (*bt.Backtest, error) {
	se := &StrategyExecution{
		StrategyName: backtest.StrategyName,
		TimeframeSec: backtest.TimeframeSec,
		Symbol:       backtest.Symbol,
		StartTime:    backtest.StartTime,
		EndTime:      backtest.EndTime,
		Status:       StrategyExecutionStatus(backtest.Status),
	}
	result := r.db.Create(se)
	if result.Error != nil {
		return nil, result.Error
	}

	backtest.Id = se.StrategyExecutionId.String()
	return &backtest, nil
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

func (r *repository) GetBacktest(id string) (*bt.Backtest, error) {
	strategyExecution, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	return strategyExecution.ToBacktest(), err
}
