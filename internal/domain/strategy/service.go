package strategy

import (
	"context"
	"stregy/internal/domain/user"
)

type Service interface {
	GetByUUID(ctx context.Context, id string) (*Strategy, error)
	Create(ctx context.Context, strategy CreateStrategyDTO, user *user.User) (*Strategy, error)
}

type service struct {
	repository Repository
	storage    Storage
}

func NewService(repository Repository, storage Storage) Service {
	return &service{repository: repository, storage: storage}
}

func (s *service) Create(ctx context.Context, dto CreateStrategyDTO, user *user.User) (strategy *Strategy, err error) {
	strategy = &Strategy{Name: dto.Name, Description: dto.Description}
	strategy, err = s.repository.Create(ctx, *strategy)
	if err != nil {
		return nil, err
	}
	s.storage.SaveStrategy(dto.Implementation, user.ID, strategy.ID)
	return strategy, nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (strategy *Strategy, err error) {
	strategy, err = s.repository.GetOne(ctx, uuid)
	if err != nil {
		return nil, err
	}

	//TODO: get strategy implementation

	return strategy, nil
}
