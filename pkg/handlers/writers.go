package handlers

import (
	"encoding/json"
	"net/http"
	"stregy/pkg/logging"
)

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
