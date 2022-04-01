package symbol

import (
	"context"
)

type Service interface {
	Exists(ctx context.Context, name string) bool
	Create(ctx context.Context, name string) (Symbol, error)
	GetAll(ctx context.Context) ([]Symbol, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Exists(ctx context.Context, name string) bool {
	return s.repository.Exists(ctx, name)
}
func (s *service) Create(ctx context.Context, name string) (Symbol, error) {
	return s.repository.Create(ctx, name)
}

func (s *service) GetAll(ctx context.Context) ([]Symbol, error) {
	return s.repository.GetAll(ctx)
}
