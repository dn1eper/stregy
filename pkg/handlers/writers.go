package handlers

import (
	"encoding/json"
	"net/http"
	"stregy/pkg/logging"
)

type Error struct {
	Message string `json:"error"`
}

func ReturnError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	if message != "" {
		resp := Error{Message: message}
		JsonResponseWriter(w, resp)
	}
}

func JsonResponseWriter(w http.ResponseWriter, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logger := logging.GetLogger()
		logger.Fatalf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	}
}
