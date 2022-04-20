package composites

import (
	position1 "stregy/internal/adapters/pgorm/position"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
)

type PositionComposite struct {
	Repository position.Repository
	Service    position.Service
}

func NewPositionComposite(composite *PGormComposite, orderService order.Service) (*PositionComposite, error) {
	repository := position1.NewRepository(composite.db)
	service := position.NewService(repository, orderService)

	return &PositionComposite{
		Repository: repository,
		Service:    service,
	}, nil
}
