package composites

import (
	"stregy/internal/domain/order"
)

type OrderComposite struct {
	Service order.Service
}

func NewOrderComposite() (*OrderComposite, error) {
	service := order.NewService()
	return &OrderComposite{
		Service: service,
	}, nil
}
