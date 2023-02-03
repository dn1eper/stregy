package symbol

import (
	"stregy/internal/domain/symbol"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) symbol.Repository {
	return &repository{db: client}
}

func (r *repository) Create(s symbol.Symbol) (*symbol.Symbol, error) {
	sDb := Symbol(s)
	if err := r.db.Create(&sDb).Error; err != nil {
		return nil, err
	}

	return sDb.ToDomain(), nil
}

func (r *repository) GetByName(name string) (*symbol.Symbol, error) {
	symbol := &Symbol{Name: name}
	if err := r.db.First(symbol).Error; err != nil {
		return nil, err
	}

	return symbol.ToDomain(), nil
}
