package api

import (
	"encoding/json"
	"net/http"
)

type coinBalanceParams struct {
	Username string
}

type counBalanceResponse struct {
	Code int
	Balance int64
}

type apiError struct{
	Code int
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	resp := apiError{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func (w http.ResponseWriter, err error) {
		writeError(w, http.StatusBadRequest, err.Error())
	}
	internalErrorHandler = func (w http.ResponseWriter) {
		writeError(w, http.StatusInternalServerError, "Internal Server Error")
	}

)