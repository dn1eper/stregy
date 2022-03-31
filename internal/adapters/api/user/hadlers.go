package user

import (
	"net/http"
	"stregy/internal/adapters/api"

	"stregy/internal/domain/user"

	"github.com/julienschmidt/httprouter"
)

const (
	userURL  = "/api/user/:user_id"
	usersURL = "/api/users"
)

type handler struct {
	userService user.Service
}

func NewHandler(service user.Service) api.Handler {
	return &handler{userService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetAllUsers)
	// TODO: add more routes
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("not implemented")
	//users := h.userService.GetAll(context.Background(), 0, 0)
	// w.Write([]byte("users"))
	// w.WriteHeader(http.StatusOK)
}

// TODO: implement more hadlers
