package backtester

import (
	"stregy/internal/domain/backtester"

	"github.com/stretchr/testify/mock"
)

type MockedRepository struct {
	mock.Mock
}

func NewMockedRepository() backtester.Repository {
	return new(MockedRepository)
}
