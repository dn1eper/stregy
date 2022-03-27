package user

import (
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) (*User, error)
	Create(ctx context.Context, dto *CreateUserDTO) (*User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, dto *CreateUserDTO) (*User, error) {
	user := &User{Name: dto.Name, Email: dto.Email, PassHash: dto.PassHash}
	return s.repository.Create(ctx, user)
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	return s.repository.GetOne(ctx, uuid)
}
