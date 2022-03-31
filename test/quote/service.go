package quote

import (
	"stregy/internal/domain/quote"
)

func NewMockedService() quote.Service {
	repository := NewMockedRepository()
	return quote.NewService(repository)
}
