package api

import (
	"net/http"
	"stregy/pkg/handlers"
)

type Error struct {
	Message string `json:"error"`
}

func ReturnError(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	if message != "" {
		resp := Error{Message: message}
		handlers.JsonResponseWriter(w, resp)
	}
}
