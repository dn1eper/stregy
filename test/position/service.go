package position

import (
	"stregy/internal/domain/position"
	"stregy/test/order"
)

func NewMockedService() position.Service {
	repository := NewMockedRepository()
	orderService := order.NewMockedService()
	return position.NewService(repository, orderService)
}
