package stratexec

import (
	"stregy/internal/domain/backtester"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) *repository {
	return &repository{db: client}
}

func (r *repository) Create(bt backtester.Backtester) (*backtester.Backtester, error) {
	se := &StrategyExecution{
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
