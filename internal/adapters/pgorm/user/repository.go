package user

import (
	"stregy/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) user.Repository {
	return &repository{db: client}
}

func (r *repository) GetOne(id string) (*user.User, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user := &User{ID: parsed}
	result := r.db.First(user)
	return user.ToDomain(), result.Error
}

func (r *repository) Create(user *user.User) (*user.User, error) {
	dbUser := &User{Name: user.Name, Email: user.Email, PassHash: user.PassHash}
	result := r.db.Create(dbUser)
	return dbUser.ToDomain(), result.Error
}

func (r *repository) GetByAPIKey(apiKey string) (*user.User, error) {
	apiUUID, err := uuid.Parse(apiKey)
	if err != nil {
		return nil, err
	}
	user := &User{APIKey: apiUUID}
	result := r.db.First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user.ToDomain(), nil
}
