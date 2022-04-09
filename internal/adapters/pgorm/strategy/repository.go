package strategy

import (
	"context"
	"stregy/internal/domain/strategy"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) strategy.Repository {
	return &repository{db: client}
}

func (r *repository) GetOne(ctx context.Context, id string) (*strategy.Strategy, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	strategy := &Strategy{ID: parsed}
	result := r.db.First(strategy)
	return strategy.ToDomain(), result.Error
}

func (r *repository) Create(ctx context.Context, strategy strategy.Strategy) (*strategy.Strategy, error) {
	db_strategy := &Strategy{Name: strategy.Name, Description: strategy.Description}
	result := r.db.Create(db_strategy)
	return db_strategy.ToDomain(), result.Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	strategy := &Strategy{ID: parsed}
	result := r.db.Delete(strategy)
	return result.Error
}
