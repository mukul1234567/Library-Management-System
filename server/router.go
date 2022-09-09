package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukul1234567/Library-Management-System/api"
)

const (
// versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {
	// v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	// router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet).Headers(versionHeader, v1)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
