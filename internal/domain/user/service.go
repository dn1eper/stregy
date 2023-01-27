package user

import (
	"crypto/sha256"
	"encoding/hex"
)

type Service interface {
	GetByUUID(id string) (*User, error)
	GetByAPIKey(apiKey string) (*User, error)
	Create(dto *CreateUserDTO) (*User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(dto *CreateUserDTO) (user *User, err error) {
	user = &User{Name: dto.Name, Email: dto.Email, PassHash: hashPassword(dto.Password)}
	return s.repository.Create(user)
}

func (s *service) GetByUUID(uuid string) (*User, error) {
	return s.repository.GetOne(uuid)
}

func (s *service) GetByAPIKey(apiKey string) (*User, error) {
	return s.repository.GetByAPIKey(apiKey)
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
