package exgaccount

import (
	"context"
	"net/http"
	"stregy/internal/adapters/api"
	userapi "stregy/internal/adapters/api/user"
	"stregy/internal/domain/user"
	"stregy/pkg/handlers"

	log "github.com/sirupsen/logrus"

	exgaccount1 "stregy/internal/domain/exgaccount"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	ExchangeAccountURL = "/api/exchange-account"
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
	authHandler := userapi.AuthenticationHandler(jsonHandler, h.userService)
	router.POST(ExchangeAccountURL, handlers.ToSimpleHandler(authHandler))
	authHandler = userapi.AuthenticationHandler(h.GetExchangeAccount, h.userService)
	router.GET(ExchangeAccountURL, handlers.ToSimpleHandler(authHandler))
}

func (h *handler) CreateExchangeAccount(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	dto := exgaccount1.CreateExchangeAccountDTO{}
	mapstructure.Decode(args["json"], &dto)
	user := user.User{}
	mapstructure.Decode(args["user"], &user)

	exchangeAccount, err := h.exgAccountService.Create(context.TODO(), dto, &user)
	if err != nil {
		log.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, "")
		return
	}

	handlers.JsonResponseWriter(w, map[string]string{"exchange_account_id": exchangeAccount.ExchangeAccountID})
}

func (h *handler) GetExchangeAccount(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	exchangeAccount, err := h.exgAccountService.GetOne(context.TODO(), r.URL.Query().Get("id"))
	if err != nil {
		log.Error(err.Error())
		handlers.ReturnError(w, http.StatusInternalServerError, "")
		return
	}
	dto := exgaccount1.GetExchangeAccountDTO{
		ConnectionString: exchangeAccount.ConnectionString,
		Name:             exchangeAccount.ExchangeAccountName,
	}
	handlers.JsonResponseWriter(w, dto)
}
