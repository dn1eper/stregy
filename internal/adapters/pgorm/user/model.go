package user

import (
	"stregy/internal/domain/user"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name     string    `gorm:"type:varchar(100);unique;not null"`
	Email    string    `gorm:"type:string;unique;not null"`
	PassHash string    `gorm:"type:text;not null"`
	APIKey   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (u *User) ToDomain() *user.User {
	return &user.User{
		ID:       u.ID.String(),
		Name:     u.Name,
		Email:    u.Email,
		PassHash: u.PassHash,
		APIKey:   u.APIKey.String(),
	}
}
