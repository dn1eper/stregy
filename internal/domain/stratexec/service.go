package stratexec

import (
	"context"
	"stregy/internal/domain/user"
)

type Service interface {
	Create(ctx context.Context, se StrategyExecution, user *user.User) (*StrategyExecution, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(ctx context.Context, se StrategyExecution, user *user.User) (*StrategyExecution, error) {
	seDB, err := s.repository.Create(ctx, se)
	if err != nil {
		return nil, err
	}
	return seDB, nil
}
