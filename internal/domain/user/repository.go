package user

import (
	"context"
)

type Repository interface {
	GetOne(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (*User, error)
	GetByAPIKey(ctx context.Context, apiKey string) (*User, error)
}
