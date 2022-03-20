package user

import (
	"context"
	"strconv"
	"stregy/internal/domain/user"

	"gorm.io/gorm"
)

type storage struct {
	db *gorm.DB
}

func NewStorage(client *gorm.DB) user.Storage {
	return &storage{db: client}
}

func (s *storage) GetOne(ctx context.Context, uuid string) (*user.User, error) {
	uuid_int, _ := strconv.ParseInt(uuid, 10, 64)
	user := &User{UUID: uuid_int}
	result := s.db.First(user)
	return user.ToDomain(), result.Error
}

func (s *storage) Create(ctx context.Context, user *user.User) (*user.User, error) {
	db_user := &User{Name: user.Name, Email: user.Email, PassHash: user.PassHash}
	result := s.db.Create(db_user)
	return db_user.ToDomain(), result.Error
}
