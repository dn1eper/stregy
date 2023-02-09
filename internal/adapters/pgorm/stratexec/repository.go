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

func (r *repository) Create(backtester bt.Backtester) (*bt.Backtester, error) {
	se := &StrategyExecution{
		StrategyName: backtester.StrategyName,
		TimeframeSec: backtester.TimeframeSec,
		SymbolName:   backtester.Symbol.Name,
		StartTime:    backtester.StartTime,
		EndTime:      backtester.EndTime,
		Status:       StrategyExecutionStatus(backtester.Status),
	}
	result := r.db.Create(se)
	if result.Error != nil {
		return nil, result.Error
	}

	backtester.ID = se.StrategyExecutionId.String()
	return &backtester, nil
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

func (r *repository) GetBacktest(id string) (*bt.Backtester, error) {
	strategyExecution, err := r.Get(id)
	if err != nil {
		return nil, err
	}
	return strategyExecution.ToBacktest(), err
}
