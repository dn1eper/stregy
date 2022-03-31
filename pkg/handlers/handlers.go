package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type JsonHandle func(http.ResponseWriter, *http.Request, httprouter.Params, interface{})

func JsonHandler(h JsonHandle, dst interface{}) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		contentType := r.Header.Get("Content-Type")

		if contentType != "" {
			mt, _, err := mime.ParseMediaType(contentType)
			if err != nil {
				http.Error(w, "Malformed Content-Type header", http.StatusBadRequest)
				return
			}

			if mt != "application/json" {
				http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&dst)
		if err != nil {
			var syntaxError *json.SyntaxError
			var unmarshalTypeError *json.UnmarshalTypeError
			var msg string

			switch {
			case errors.As(err, &syntaxError):
				msg = fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)

			case errors.Is(err, io.ErrUnexpectedEOF):
				msg = "Request body contains badly-formed JSON"

			case errors.As(err, &unmarshalTypeError):
				msg = fmt.Sprintf("Request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)

			case strings.HasPrefix(err.Error(), "json: unknown field "):
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
				msg = fmt.Sprintf("Request body contains unknown field %s", fieldName)

			case errors.Is(err, io.EOF):
				msg = "Request body must not be empty"

			case err.Error() == "http: request body too large":
				msg = "Request body must not be larger than 1MB"

			default:
				msg = err.Error()
			}
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		err = dec.Decode(&struct{}{})
		if err != io.EOF {
			http.Error(w, "Request body must only contain a single JSON object", http.StatusBadRequest)
			return
		}

		h(w, r, ps, dst)
	}
}
