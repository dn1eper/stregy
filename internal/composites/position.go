package composites

import (
	position1 "stregy/internal/adapters/pgorm/position"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
)

type PositionComposite struct {
	Service position.Service
}

func NewPositionComposite(repoComposite *PGormComposite, orderService order.Service) (*PositionComposite, error) {
	repository := position1.NewRepository(repoComposite.db)
	service := position.NewService(repository, orderService)
	return &PositionComposite{
		Service: service,
	}, nil
}
