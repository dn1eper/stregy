package user

import (
	"context"
)

type Storage interface {
	GetOne(ctx context.Context, uuid string) (*User, error)
	GetAll(ctx context.Context, limit, offset int) ([]*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, user *User) error
}
