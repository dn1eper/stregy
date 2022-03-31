package composites

import (
	"stregy/internal/adapters/api"
	strategy1 "stregy/internal/adapters/api/strategy"
	strategy2 "stregy/internal/adapters/pgorm/strategy"
	"stregy/internal/domain/strategy"
)

type StrategyComposite struct {
	Repository strategy.Repository
	Service    strategy.Service
	Handler    api.Handler
}

func NewStrategyComposite(composite *PGormComposite) (*StrategyComposite, error) {
	repository := strategy2.NewRepository(composite.db)
	service := strategy.NewService(repository)
	handler := strategy1.NewHandler(service)
	return &StrategyComposite{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}, nil
}
