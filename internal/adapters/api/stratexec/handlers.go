package stratexec

import (
	"context"
	"net/http"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"

	stratexec1 "stregy/internal/domain/stratexec"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	strategyExecutionURL = "/api/strategy-execution"
)

type handler struct {
	strategyExecutionService stratexec1.Service
	userService              user.Service
}

func NewHandler(service stratexec1.Service, userService user.Service) api.Handler {
	return &handler{strategyExecutionService: service, userService: userService}
}

func (h *handler) Register(router *httprouter.Router) {
	createSEHandler := handlers.JsonHandler(h.CreateStrategyExecution, &stratexec1.CreateStrategyExecutionDTO{})
	userHandler := handlers.ToSimpleHandler(userapi.APIKeyHandler(createSEHandler, h.userService))
	router.POST(strategyExecutionURL, userHandler)
}

func (h *handler) CreateStrategyExecution(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	logger := logging.GetLogger()

	dto := stratexec1.CreateStrategyExecutionDTO{}
	mapstructure.Decode(args["json"], &dto)
	user := user.User{}
	mapstructure.Decode(args["user"], &user)

	se, err := h.strategyExecutionService.Create(context.TODO(), dto, &user)
	if err != nil {
		logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		resp := api.Error{Message: err.Error()}
		handlers.JsonResponseWriter(w, resp)
		return
	}

	handlers.JsonResponseWriter(w, map[string]string{"strategy_execution_id": se.ID})
}
