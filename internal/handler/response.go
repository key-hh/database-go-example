package handler

import (
	"encoding/json"
	"net/http"
)

func writeResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(resp)
}

func writeErrorResponse(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	resp := struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	}
	json.NewEncoder(w).Encode(resp)
}
