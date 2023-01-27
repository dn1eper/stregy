package backtester

import (
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/tick"
)

type Service interface {
	Run(b *Backtester) error
	Create(dto BacktesterDTO) (*Backtester, error)
}

type service struct {
	repository      Repository
	tickService     tick.Service
	quoteService    quote.Service
	exgAccService   exgaccount.Service
	strategyService strategy.Service
	positionService position.Service

	stratexecProjectPath string

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

func (s *service) Create(dto BacktesterDTO) (*Backtester, error) {
	bt := Backtester{
		StrategyName:        dto.StrategyName,
		StartDate:           dto.StartDate,
		EndDate:             dto.EndDate,
		Symbol:              dto.Symbol,
		Timeframe:           dto.Timeframe,
		HighOrderResolution: dto.HighOrderResolution,
		Status:              Created,
	}
	return s.repository.Create(bt)
}

// Update Quotes in sync, mb prep new Quotes in seperate goroutine
func (s *service) Run(bt *Backtester) (err error) {
	return err
}
