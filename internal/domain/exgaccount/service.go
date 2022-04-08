package exgaccount

import (
	"context"
	"stregy/internal/domain/user"
)

type Service interface {
	Create(ctx context.Context, dto CreateExchangeAccountDTO, user *user.User) (*ExchangeAccount, error)
	GetAll(ctx context.Context, userID string) ([]*ExchangeAccount, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, dto CreateExchangeAccountDTO, user *user.User) (*ExchangeAccount, error) {
	exgAccount := &ExchangeAccount{
		ExchangeID:          dto.ExchangeID,
		ConnectionString:    dto.ConnectionString,
		ExchangeAccountName: dto.Name,
	}

	strategy, err := s.repository.Create(ctx, *exgAccount, user)
	if err != nil {
		return nil, err
	}

	return strategy, nil
}

func (s *service) GetAll(ctx context.Context, userID string) ([]*ExchangeAccount, error) {
	return s.repository.GetAll(ctx, userID)
}
