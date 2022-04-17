package composites

import (
	stratexec1 "stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/domain/stratexec"
)

type StrategyExecutionComposite struct {
	Repository stratexec.Repository
	Service    stratexec.Service
}

func NewStrategyExecutionComposite(
	composite *PGormComposite,
) (*StrategyExecutionComposite, error) {
	repository := stratexec1.NewRepository(composite.db)
	service := stratexec.NewService(repository)
	return &StrategyExecutionComposite{
		Repository: repository,
		Service:    service,
	}, nil
}
