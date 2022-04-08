package strategy

import (
	"context"
	"net/http"
	"stregy/internal/adapters/api"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"

	strategy1 "stregy/internal/domain/strategy"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	strategyURL = "/api/strategy"
)

type handler struct {
	strategyService strategy1.Service
}

func NewHandler(service strategy1.Service) api.Handler {
	return &handler{strategyService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	createStrategyHandler := handlers.ToSimpleHandler(
		handlers.JsonHandler(h.CreateStrategy, &strategy1.CreateStrategyDTO{}))
	router.POST(strategyURL, createStrategyHandler)
}

func (h *handler) CreateStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	logger := logging.GetLogger()

	dto := strategy1.CreateStrategyDTO{}
	mapstructure.Decode(args["json"], &dto)
	user := user.User{}
	mapstructure.Decode(args["user"], &user)

	strat, err := h.strategyService.Create(context.TODO(), dto, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
	}

	handlers.JsonResponseWriter(w, map[string]string{"strategy_id": strat.ID})
}
