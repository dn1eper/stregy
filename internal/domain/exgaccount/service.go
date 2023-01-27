package exgaccount

import (
	"stregy/internal/domain/user"
)

type Service interface {
	Create(dto CreateExchangeAccountDTO, user *user.User) (*ExchangeAccount, error)
	GetAll(userID string) ([]*ExchangeAccount, error)
	GetOne(exgAccountID string) (*ExchangeAccount, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(dto CreateExchangeAccountDTO, user *user.User) (*ExchangeAccount, error) {
	exgAccount := &ExchangeAccount{
		ExchangeID:          dto.ExchangeID,
		ConnectionString:    dto.ConnectionString,
		ExchangeAccountName: dto.Name,
	}

	strategy, err := s.repository.Create(*exgAccount, user)
	if err != nil {
		return nil, err
	}

	return strategy, nil
}

func (s *service) GetAll(userID string) ([]*ExchangeAccount, error) {
	return s.repository.GetAll(userID)
}

func (s *service) GetOne(exgAccountID string) (*ExchangeAccount, error) {
	return s.repository.GetOne(exgAccountID)
}
