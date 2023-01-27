package exgaccount

import (
	"stregy/internal/domain/user"
)

type Repository interface {
	GetAll(userID string) ([]*ExchangeAccount, error)
	GetOne(exgAccountID string) (*ExchangeAccount, error)
	Create(exgAccount ExchangeAccount, user *user.User) (*ExchangeAccount, error)
}
