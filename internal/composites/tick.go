package composites

import (
	tick1 "stregy/internal/adapters/pgorm/tick"
	"stregy/internal/domain/tick"
)

type TickComposite struct {
	Repository tick.Repository
	Service    tick.Service
}

func NewTickComposite(composite *PGormComposite) (*TickComposite, error) {
	repository := tick1.NewRepository(composite.db)
	service := tick.NewService(repository)

	return &TickComposite{
		Repository: repository,
		Service:    service,
	}, nil
}
