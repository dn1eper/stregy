package backtester

import (
	"context"
	"errors"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"time"
)

type Service interface {
	Run(ctx context.Context, b *Backtester) error
	Create(ctx context.Context, dto BacktesterDTO, userID string) (*Backtester, error)
}

type service struct {
	repository      Repository
	quoteService    quote.Service
	exgAccService   exgaccount.Service
	strategyService strategy.Service
	positionService position.Service

	positions []*position.Position
}

func NewService(
	repository Repository,
	quoteService quote.Service,
	exgAccService exgaccount.Service,
	positionService position.Service,
	strategyService strategy.Service,
) Service {
	return &service{
		repository:      repository,
		quoteService:    quoteService,
		exgAccService:   exgAccService,
		strategyService: strategyService,
		positionService: positionService,
	}
}

func (s *service) Create(ctx context.Context, dto BacktesterDTO, userID string) (*Backtester, error) {
	exgAccount, err := s.exgAccService.GetOne(context.TODO(), dto.ExchangeAccountID)
	if err != nil {
		return nil, err
	}
	if userID != exgAccount.UserID {
		return nil, errors.New("incorrect exchange account id")
	}

	startDate, _ := time.Parse("2006-01-02", dto.StartDate)
	endDate, _ := time.Parse("2006-01-02", dto.EndDate)
	strat, err := s.strategyService.GetByUUID(context.TODO(), dto.StrategyID)
	bt := Backtester{
		Strategy:  *strat,
		StartDate: startDate,
		EndDate:   endDate,
		Symbol:    dto.Symbol,
		Timeframe: dto.Timeframe,
		Status:    Created,
	}
	return s.repository.CreateBacktester(ctx, bt, dto.ExchangeAccountID)
}

func (s *service) Run(ctx context.Context, b *Backtester) (err error) {
	var quotes []quote.Quote

	// TODO: get quotes

	for _, q := range quotes {
		b.Strategy.Implementation.OnQuote(ctx, q)

		for _, p := range s.positions {
			if p.Status == position.Created && p.MainOrder.IsTouched(q) {
				p, err = s.positionService.Open(ctx, *p, p.MainOrder.Size)
				if err != nil {
					return err
				}

				b.Strategy.Implementation.OnOrder(ctx, p.MainOrder)
				b.Strategy.Implementation.OnPosition(ctx, *p)

			} else if p.Status == position.Open {
				if p.TakeOrder.IsTouched(q) {
					p, err = s.positionService.TakeProfit(ctx, *p, p.MainOrder.Size)
					if err != nil {
						return err
					}

					b.Strategy.Implementation.OnOrder(ctx, p.TakeOrder)
					b.Strategy.Implementation.OnPosition(ctx, *p)

				} else if p.StopOrder.IsTouched(q) {
					p, err = s.positionService.StopLoss(ctx, *p, p.MainOrder.Size)
					if err != nil {
						return err
					}

					b.Strategy.Implementation.OnOrder(ctx, p.TakeOrder)
					b.Strategy.Implementation.OnPosition(ctx, *p)
				}
			}
		}
	}

	return nil
}
