package composites

import (
	order1 "stregy/internal/adapters/pgorm/order"
	"stregy/internal/domain/order"
)

type OrderComposite struct {
	Service order.Service
}

func NewOrderComposite(composite *PGormComposite) (*OrderComposite, error) {
	repository := order1.NewRepository(composite.db)
	service := order.NewService(repository)
	return &OrderComposite{
		Service: service,
	}, nil
}
