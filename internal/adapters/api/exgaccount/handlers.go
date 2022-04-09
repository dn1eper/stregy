package exgaccount

import (
	"context"
	"net/http"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"
	"stregy/pkg/logging"

	exgaccount1 "stregy/internal/domain/exgaccount"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	createExchangeAccountURL = "/api/create-exchange-account"
)

type handler struct {
	userService       user.Service
	exgAccountService exgaccount1.Service
}

func NewHandler(service exgaccount1.Service, userService user.Service) api.Handler {
	return &handler{exgAccountService: service, userService: userService}
}

func (h *handler) Register(router *httprouter.Router) {
	jsonHandler := handlers.JsonHandler(h.CreateExchangeAccount, &exgaccount1.CreateExchangeAccountDTO{})
	userHandler := handlers.ToSimpleHandler(userapi.APIKeyHandler(jsonHandler, h.userService))
	router.POST(createExchangeAccountURL, userHandler)
}

func (h *handler) CreateExchangeAccount(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	logger := logging.GetLogger()

	dto := exgaccount1.CreateExchangeAccountDTO{}
	mapstructure.Decode(args["json"], &dto)
	user := user.User{}
	mapstructure.Decode(args["user"], &user)

	exchangeAccount, err := h.exgAccountService.Create(context.TODO(), dto, &user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error(err.Error())
	}

	handlers.JsonResponseWriter(w, map[string]string{"exchange_account_id": exchangeAccount.ExchangeAccountID})
}
