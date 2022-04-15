package composites

import (
	"stregy/internal/adapters/api"
	stratexec1 "stregy/internal/adapters/api/stratexec"
	stratexec2 "stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/stratexec"
	"stregy/internal/domain/user"
)

type StrategyExecutionComposite struct {
	Repository stratexec.Repository
	Service    stratexec.Service
	Handler    api.Handler
}

func NewStrategyExecutionComposite(composite *PGormComposite, userService user.Service, exgAccService exgaccount.Service) (*StrategyExecutionComposite, error) {
	repository := stratexec2.NewRepository(composite.db)
	service := stratexec.NewService(repository, exgAccService)
	handler := stratexec1.NewHandler(service, userService)
	return &StrategyExecutionComposite{
		Repository: repository,
		Service:    service,
		Handler:    handler,
	}, nil
}
