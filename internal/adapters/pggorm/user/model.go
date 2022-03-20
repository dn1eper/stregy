package user

import (
	"stregy/internal/domain/user"
)

type User struct {
	UUID     int64  `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);unique;not null"`
	Email    string `gorm:"type:string;unique;not null"`
	PassHash string `gorm:"type:text;not null"`
}

func (u *User) ToDomain() *user.User {
	return &user.User{UUID: string(u.UUID), Name: u.Name, Email: u.Email, PassHash: u.PassHash}
}
