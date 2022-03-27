package user

import (
	"context"
	"strconv"
	"stregy/internal/domain/user"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) user.Repository {
	return &repository{db: client}
}

func (r *repository) GetOne(ctx context.Context, uuid string) (*user.User, error) {
	uuid_int, _ := strconv.ParseInt(uuid, 10, 64)
	user := &User{UUID: uuid_int}
	result := r.db.First(user)
	return user.ToDomain(), result.Error
}

func (r *repository) Create(ctx context.Context, user *user.User) (*user.User, error) {
	db_user := &User{Name: user.Name, Email: user.Email, PassHash: user.PassHash}
	result := r.db.Create(db_user)
	return db_user.ToDomain(), result.Error
}
