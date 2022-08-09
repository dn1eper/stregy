package backtester

import (
	"context"
	"fmt"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/tick"
)

type Service interface {
	Run(ctx context.Context, b *Backtester) error
	Create(ctx context.Context, dto BacktesterDTO) (*Backtester, error)
}

type service struct {
	repository      Repository
	tickService     tick.Service
	quoteService    quote.Service
	exgAccService   exgaccount.Service
	strategyService strategy.Service
	positionService position.Service

	positions []*position.Position
}

func NewService(
	repository Repository,
	tickService tick.Service,
	quoteService quote.Service,
	exgAccService exgaccount.Service,
	positionService position.Service,
	strategyService strategy.Service,
) Service {
	return &service{
		repository:      repository,
		tickService:     tickService,
		quoteService:    quoteService,
		exgAccService:   exgAccService,
		strategyService: strategyService,
		positionService: positionService,
	}
}

func (s *service) Create(ctx context.Context, dto BacktesterDTO) (*Backtester, error) {
	strat := strategy.Strategy{ID: dto.StrategyID}
	bt := Backtester{
		Strategy:            strat,
		StartDate:           dto.StartDate,
		EndDate:             dto.EndDate,
		Symbol:              dto.Symbol,
		Timeframe:           dto.Timeframe,
		HighOrderResolution: dto.HighOrderResolution,
		Status:              Created,
	}
	return s.repository.Create(ctx, bt)
}

// Update Quotes in sync, mb prep new Quotes in seperate goroutine
func (s *service) Run(ctx context.Context, bt *Backtester) (err error) {
	return fmt.Errorf("not implemented")
}
