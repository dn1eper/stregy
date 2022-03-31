package quote

import (
	"context"
	"time"
)

type Service interface {
	GetByInterval(ctx context.Context, symbol string, start, end time.Time, offset, pageSize int) ([]Quote, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetByInterval(ctx context.Context, symbol string, start, end time.Time, offset, pageSize int) ([]Quote, error) {
	return s.repository.GetByInterval(ctx, symbol, start, end, offset, pageSize)
}
