package user

import (
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) (*User, error)
	Create(ctx context.Context, dto *CreateUserDTO) (*User, error)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateUserDTO) (*User, error) {
	user := &User{Name: dto.Name, Email: dto.Email, PassHash: dto.PassHash}
	return s.storage.Create(ctx, user)
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	return s.storage.GetOne(ctx, uuid)
}
