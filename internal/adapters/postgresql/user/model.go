package user

import (
	"stregy/internal/domain/user"
)

type User struct {
	ID    string
	Name  string
	Email string
}

func (u *User) ToDomain() user.User {
	panic("not implemented")
}
