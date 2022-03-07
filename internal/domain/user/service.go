package user

import (
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) (*User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*User, error)
	Create(ctx context.Context, dto *CreateUserDTO) (*User, error)
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Create(ctx context.Context, dto *CreateUserDTO) (*User, error) {
	panic("not implemented")
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	return s.storage.GetOne(ctx, uuid)
}

func (s *service) GetAll(ctx context.Context, limit, offset int) ([]*User, error) {
	return s.storage.GetAll(ctx, limit, offset)
}
