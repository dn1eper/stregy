package backtester

import (
	"context"
	"stregy/internal/domain/quote"
)

type Service interface {
	Run(ctx context.Context, b *Backtester) error
}

type service struct {
	repository   Repository
	quoteService quote.Service
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Run(ctx context.Context, b *Backtester) error {
	// quotes, err := s.quoteService.GetByInterval(ctx, b.Symbol, b.StartDate, b.EndDate)
	// if err != nil {

	// }
	// execute strategy on each tic
	return nil
}
