package composites

import (
	"stregy/internal/adapters/api"
	backtester1 "stregy/internal/adapters/api/backtester"
	stratexec1 "stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/domain/backtester"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/tick"
	"stregy/internal/domain/user"
)

type BacktesterComposite struct {
	Service backtester.Service
	Handler api.Handler
}

func NewBacktesterComposite(
	pgormComposite *PGormComposite,
	exgAccService exgaccount.Service,
	strategyService strategy.Service,
	userService user.Service,
	tickService tick.Service,
	quoteService quote.Service,
	positionService position.Service,
	executor backtester.Executor,
) (*BacktesterComposite, error) {
	repository := stratexec1.NewRepository(pgormComposite.db)
	service := backtester.NewService(
		repository, tickService, quoteService, exgAccService,
		positionService, strategyService, executor)
	handler := backtester1.NewHandler(service, userService)
	return &BacktesterComposite{
		Service: service,
		Handler: handler,
	}, nil
}
