package composites

import (
	"stregy/internal/adapters/api"
	btapi "stregy/internal/adapters/api/bt"
	"stregy/internal/adapters/pgorm/stratexec"
	"stregy/internal/domain/btservice"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/tick"
	"stregy/internal/domain/user"
)

type BacktesterComposite struct {
	Service btservice.Service
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
) (*BacktesterComposite, error) {
	repository := stratexec.NewRepository(pgormComposite.db)
	service := btservice.NewService(repository, tickService, quoteService, exgAccService, positionService)
	handler := btapi.NewHandler(service, userService)
	return &BacktesterComposite{
		Service: service,
		Handler: handler,
	}, nil
}
