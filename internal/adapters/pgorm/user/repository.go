package user

import (
	"context"
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

func (r *repository) GetOne(ctx context.Context, id string) (*user.User, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	user := &User{ID: parsed}
	result := r.db.First(user)
	return user.ToDomain(), result.Error
}

func (r *repository) Create(ctx context.Context, user *user.User) (*user.User, error) {
	dbUser := &User{Name: user.Name, Email: user.Email, PassHash: user.PassHash}
	result := r.db.Create(dbUser)
	return dbUser.ToDomain(), result.Error
}
