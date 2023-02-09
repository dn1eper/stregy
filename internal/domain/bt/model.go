package bt

import (
	"stregy/internal/domain/broker"
	"stregy/internal/domain/order"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/symbol"
	"time"
)

type StrategyExecutionStatus string

const (
	Created    StrategyExecutionStatus = "Created"
	Running    StrategyExecutionStatus = "Running"
	Finished   StrategyExecutionStatus = "Finished"
	Terminated StrategyExecutionStatus = "Terminated"
	Crashed    StrategyExecutionStatus = "Crashed"
)

type Backtester struct {
	ID                    string
	StrategyName          string
	StartTime             time.Time
	EndTime               time.Time
	Symbol                symbol.Symbol
	TimeframeSec          int
	Status                StrategyExecutionStatus
	AccountHistoryService AccountHistoryReport

	logger broker.Logger

	strategy strategy.Strategy
	balance  float64

	curTime   time.Time
	lastPrice float64

	termChan chan bool

	orders        map[int64]*order.Order
	positions     map[int64]*order.Position
	orderCount    int64
	positionCount int64
	orderHistory  []*order.Order
}

var _ broker.Broker = (*Backtester)(nil)
