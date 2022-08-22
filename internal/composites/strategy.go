package composites

import (
	"stregy/internal/adapters/api"
	strategy1 "stregy/internal/adapters/api/strategy"
	strategy3 "stregy/internal/adapters/fs/strategy"
	strategy2 "stregy/internal/adapters/pgorm/strategy"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/user"
)

type StrategyComposite struct {
		  strategy.Repository
	Service     strategy.Service
	userService user.Service
	Handler     api.Handler
}

func NewStrategyComposite(composite *PGormComposite, userService user.Service) (*StrategyComposite, error) {
	storage := strategy3.NewStorage()
	repository := strategy2.NewRepository(composite.db)
	service := strategy.NewService(repository, storage)
	handler := strategy1.NewHandler(service, userService)
	return &StrategyComposite{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}, nil
}
