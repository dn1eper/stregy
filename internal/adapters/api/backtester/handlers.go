package backtester

import (
	"context"
	"net/http"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/domain/backtester"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	strategyExecutionURL = "/api/backtest"
)

type handler struct {
	backtesterService backtester.Service
	exgAccService     exgaccount.Service
	userService       user.Service
}

func NewHandler(
	backtesterService backtester.Service,
	userService user.Service,
) api.Handler {
	return &handler{
		backtesterService: backtesterService,
		userService:       userService,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	createSEHandler := handlers.JsonHandler(h.ExecuteBacktestHandler, &backtester.BacktesterDTO{})
	userHandler := userapi.AuthenticationHandler(createSEHandler, h.userService)
	router.POST(strategyExecutionURL, handlers.ToSimpleHandler(userHandler))
}

func (h *handler) ExecuteBacktestHandler(
	w http.ResponseWriter,
	r *http.Request,
	params httprouter.Params,
	args map[string]interface{},
) {
	logger := logging.GetLogger()
	// Parse and validate request.
	user := user.User{}
	mapstructure.Decode(args["user"], &user)
	dto := backtester.BacktesterDTO{}
	mapstructure.Decode(args["json"], &dto)

	// Create db record.
	backtesterDB, err := h.backtesterService.Create(context.TODO(), dto, user.ID)
	if err != nil {
		logger.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, "")
		return
	}

	// Execute.
	err = h.backtesterService.Run(context.TODO(), backtesterDB)
	if err != nil {
		logger.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, err.Error())
		return
	}

	handlers.JsonResponseWriter(w, map[string]string{"strategy_execution_id": backtesterDB.ID})
}
