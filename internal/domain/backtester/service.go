package backtester

import (
	"context"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
)

type Service interface {
	Run(ctx context.Context, b *Backtester) error
}

type service struct {
	repository      Repository
	quoteService    quote.Service
	positionService position.Service

	positions []*position.Position
}

func NewService(repository Repository, quoteService quote.Service) Service {
	return &service{repository: repository, quoteService: quoteService}
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
