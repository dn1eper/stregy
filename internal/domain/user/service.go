package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
)

type Service interface {
	GetByUUID(ctx context.Context, id string) (*User, error)
	GetByAPIKey(ctx context.Context, apiKey string) (*User, error)
	Create(ctx context.Context, dto *CreateUserDTO) (*User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, dto *CreateUserDTO) (user *User, err error) {
	user = &User{Name: dto.Name, Email: dto.Email, PassHash: hashPassword(dto.Password)}
	return s.repository.Create(ctx, user)
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (*User, error) {
	return s.repository.GetOne(ctx, uuid)
}

func (s *service) GetByAPIKey(ctx context.Context, apiKey string) (*User, error) {
	return s.repository.GetByAPIKey(ctx, apiKey)
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
