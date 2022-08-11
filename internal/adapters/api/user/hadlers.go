package user

import (
	"context"
	"net/http"
	"stregy/internal/adapters/api"
	"stregy/pkg/handlers"

	log "github.com/sirupsen/logrus"

	"stregy/internal/domain/user"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchellh/mapstructure"
)

const (
	createUserURL = "/api/create-user"
)

type UserHandle func(http.ResponseWriter, *http.Request, httprouter.Params, *user.User)

type handler struct {
	userService user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	jsonHandler := handlers.JsonHandler(h.CreateUser, user.CreateUserDTO{})
	handler := handlers.ToSimpleHandler(jsonHandler)
	router.POST(createUserURL, handler)
	// TODO: add more routes

}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params, args map[string]interface{}) {
	dto := user.CreateUserDTO{}
	mapstructure.Decode(args["json"], &dto)

	user, err := h.userService.Create(context.TODO(), &dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error(err.Error())
		return
	}

	handlers.JsonResponseWriter(w, map[string]string{"user_id": user.ID, "api_key": user.APIKey})
}

func AuthenticationHandler(h handlers.HandleWithArgs, s user.Service) handlers.HandleWithArgs {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params, args map[string]interface{}) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			http.Error(w, "X-API-Key not provided in headers", http.StatusBadRequest)
			return
		}

		u, err := s.GetByAPIKey(context.TODO(), apiKey)
		if err != nil {
			http.Error(w, "Invalid API key", http.StatusNotFound)
			return
		}
		args["user"] = u

		h(w, r, ps, args)
	}
}
