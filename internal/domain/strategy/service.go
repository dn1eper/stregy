package strategy

import "context"

type Service interface {
	GetByUUID(ctx context.Context, uuid string) (*Strategy, error)
	Create(ctx context.Context, strategy *CreateStrategyDTO) (*Strategy, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, dto *CreateStrategyDTO) (*Strategy, error) {
	strategy := &Strategy{Name: dto.Name, Description: dto.Description}

	_, err := s.repository.Create(ctx, strategy)
	if err != nil {
		return nil, err
	}

	//TODO: save strategy implementation

	return strategy, nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (*Strategy, error) {
	strategy, err := s.repository.GetOne(ctx, uuid)
	if err != nil {

		return nil, err
	}

	//TODO: get strategy implementation

	return strategy, nil
}
