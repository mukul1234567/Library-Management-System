package api

import (
	"encoding/json"
	"net/http"

	"github.com/mukul1234567/Library-Management-System/app"
)

type Response struct {
	Message string `json:"message"`
}
type Response1 struct {
	Message error `json:"message"`
}

func Error(rw http.ResponseWriter, status int, response interface{}) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		app.GetLogger().Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}

func Success(rw http.ResponseWriter, status int, response interface{}) {
	respBytes, err := json.Marshal(response)
	if err != nil {
		app.GetLogger().Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
