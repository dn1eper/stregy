package user

import (
	"context"
	"encoding/json"
	"net/http"
	"stregy/internal/adapters/api"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"

	"stregy/internal/domain/strategy"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	strategyURL = "/api/strategy"
)

type handler struct {
	strategyService strategy.Service
}

func NewHandler(service strategy.Service) api.Handler {
	return &handler{strategyService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	createStrategyHandler := handlers.JsonHandler(h.CreateStrategy, &strategy.CreateStrategyDTO{})
	router.POST(strategyURL, createStrategyHandler)
}

func (h *handler) CreateStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params, obj interface{}) {
	logger := logging.GetLogger()

	dto := strategy.CreateStrategyDTO{}
	mapstructure.Decode(obj, &dto)

	strat, err := h.strategyService.Create(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
	}
	JsonResponseWriter(w, map[string]string{"strategy_id": strat.ID})
}

func JsonResponseWriter(w http.ResponseWriter, resp interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logger := logging.GetLogger()
		logger.Fatalf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(jsonResp)
		w.WriteHeader(http.StatusOK)
	}
}
