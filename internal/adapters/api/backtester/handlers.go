package stratexec

import (
	"context"
	"net/http"
	"strconv"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/domain/backtester"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/stratexec"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	strategyExecutionURL = "/api/backtest"
)

type handler struct {
	backtesterService backtester.Service
	stratExecService  stratexec.Service
	exgAccService     exgaccount.Service
	strategyService   strategy.Service
	userService       user.Service
}

func NewHandler(
	backtesterService backtester.Service,
	stratExecService stratexec.Service,
	exgAccService exgaccount.Service,
	strategyService strategy.Service,
	userService user.Service,
) api.Handler {
	return &handler{
		backtesterService: backtesterService,
		stratExecService:  stratExecService,
		exgAccService:     exgAccService,
		strategyService:   strategyService,
		userService:       userService,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	createSEHandler := handlers.JsonHandler(h.Backtest, &backtester.BacktestDTO{})
	userHandler := handlers.ToSimpleHandler(userapi.APIKeyHandler(createSEHandler, h.userService))
	router.POST(strategyExecutionURL, userHandler)
}

func (h *handler) Backtest(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
	args map[string]interface{},
) {
	logger := logging.GetLogger()
	// Parse and validate request.
	user := user.User{}
	mapstructure.Decode(args["user"], &user)
	dto := backtester.BacktestDTO{}
	mapstructure.Decode(args["json"], &dto)

	userID := h.exgAccService.GetUserID(context.TODO(), dto.ExchangeAccountID)
	if userID != user.ID {
		logger.Error("incorrect exchange account id")
		api.ReturnError(w, "incorrect exchange account id")
		return
	}

	// Create db record.
	timeframe, _ := strconv.Atoi(dto.Timeframe)
	startDate, _ := time.Parse("2006-01-02", dto.StartDate)
	endDate, _ := time.Parse("2006-01-02", dto.EndDate)
	se := &stratexec.StrategyExecution{
		StrategyID:        dto.StrategyID,
		ExchangeAccountID: dto.ExchangeAccountID,
		Timeframe:         timeframe,
		Symbol:            dto.Symbol,
		StartDate:         startDate,
		EndDate:           endDate,
		Status:            stratexec.Created,
	}
	se, err := h.stratExecService.Create(context.TODO(), *se, &user)
	if err != nil {
		logger.Error(err.Error())
		api.ReturnError(w, "")
		return
	}

	// Execute.
	strat, err := h.strategyService.GetByUUID(context.TODO(), dto.StrategyID)
	backtester := backtester.Backtester{
		Strategy:  *strat,
		StartDate: startDate,
		EndDate:   endDate,
		Symbol:    dto.Symbol,
	}
	err = h.backtesterService.Run(context.TODO(), &backtester)
	if err != nil {
		logger.Error(err.Error())
		api.ReturnError(w, err.Error())
		return
	}

	handlers.JsonResponseWriter(w, map[string]string{"strategy_execution_id": se.ID})
}
