package strategy

import (
	"stregy/internal/domain/strategy"

	"github.com/google/uuid"
)

type Strategy struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:varchar()"`
}

func (u *Strategy) ToDomain() *strategy.Strategy {
	return &strategy.Strategy{
		ID: u.ID.String(),
	}
}
