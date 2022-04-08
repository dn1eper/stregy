package exgaccount

import (
	"context"
	"stregy/internal/domain/user"
)

type Repository interface {
	GetAll(ctx context.Context, userID string) ([]*ExchangeAccount, error)
	Create(ctx context.Context, exgAccount ExchangeAccount, user *user.User) (*ExchangeAccount, error)
}
