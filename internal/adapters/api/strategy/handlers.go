package strategy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"stregy/internal/adapters/api"
	user1 "stregy/internal/adapters/api/user"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"

	strategy1 "stregy/internal/domain/strategy"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
	"github.com/tjarratt/babble"
)

const (
	strategyURL       = "/api/strategy"
	strategyExistsURL = "/api/strategy-exists"
)

type handler struct {
	strategyService strategy1.Service
	userService     user.Service
}

func NewHandler(service strategy1.Service, userService user.Service) api.Handler {
	return &handler{strategyService: service, userService: userService}
}

func (h *handler) Register(router *httprouter.Router) {
	authHandler := user1.AuthenticationHandler(h.CreateStrategy, h.userService)
	router.POST(strategyURL, handlers.ToSimpleHandler(authHandler))
	router.GET(strategyExistsURL, h.StrategyExists)
}

func (h *handler) CreateStrategy(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	logger := logging.GetLogger()

	err := r.ParseMultipartForm(2000000000)
	if err != nil {
		fmt.Printf("Error parsing form: %v\n", err)
	}

	dto := strategy1.CreateStrategyDTO{}
	err = json.Unmarshal([]byte(r.Form["json"][0]), &dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
	}
	if dto.Name == "" {
		babbler := babble.NewBabbler()
		dto.Name = babbler.Babble()
		fmt.Println(dto.Name)
	}

	dto.Implementation = &(r.PostForm["file"][0])
	if len(*dto.Implementation) == 0 {
		logger.Error(err.Error())
		handlers.ReturnError(w, http.StatusBadRequest, "implementation is empty")
	}

	user := user.User{}
	mapstructure.Decode(args["user"], &user)

	strat, err := h.strategyService.Create(context.TODO(), dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
	}

	handlers.JsonResponseWriter(w, map[string]string{"strategy_id": strat.ID})
}

func (h *handler) StrategyExists(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := "true"
	_, err := h.strategyService.GetByUUID(context.TODO(), r.URL.Query().Get("id"))
	if err != nil {
		response = "false"
		fmt.Println(err.Error())
	}
	handlers.JsonResponseWriter(w, map[string]string{"response": response})
}
