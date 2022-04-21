package strategy

import (
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, id string) (*Strategy, error)
	Create(ctx context.Context, dto CreateStrategyDTO) (*Strategy, error)
}

type service struct {
	repository Repository
	storage    Storage
}

func NewService(repository Repository, storage Storage) Service {
	return &service{repository: repository, storage: storage}
}

func (s *service) Create(ctx context.Context, dto CreateStrategyDTO) (*Strategy, error) {
	strategy := &Strategy{Name: dto.Name, Description: dto.Description}
	strategy, err := s.repository.Create(ctx, *strategy)
	if err != nil {
		return nil, err
	}
	s.storage.SaveStrategy(dto.Implementation, strategy.ID)
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
