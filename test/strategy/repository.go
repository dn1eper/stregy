package strategy

import (
	"context"
	"stregy/internal/domain/strategy"

	"github.com/stretchr/testify/mock"
)

type MockedRepository struct {
	mock.Mock
}

func NewMockedRepository() strategy.Repository {
	return new(MockedRepository)
}

func (r *MockedRepository) GetOne(ctx context.Context, uuid string) (*strategy.Strategy, error) {
	panic("not implemented")
}

func (r *MockedRepository) Delete(ctx context.Context, uuid string) error {
	panic("not implemented")
}

func (r *MockedRepository) Create(ctx context.Context, strategy strategy.Strategy) (*strategy.Strategy, error) {
	panic("not implemented")
}
