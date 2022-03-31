package quote

import (
	"context"
	"stregy/internal/domain/quote"
	"time"

	"github.com/stretchr/testify/mock"
)

type MockedRepository struct {
	mock.Mock
}

func NewMockedRepository() quote.Repository {
	return new(MockedRepository)
}

func (r *MockedRepository) GetByInterval(ctx context.Context, symbol string, startTime, endTime time.Time, offset, pageSize int) ([]quote.Quote, error) {
	panic("not implemented")
}
