package position

import (
	"stregy/internal/domain/position"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) position.Repository {
	return &repository{db: client}
}
