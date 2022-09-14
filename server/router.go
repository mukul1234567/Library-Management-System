package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukul1234567/Library-Management-System/api"
	"github.com/mukul1234567/Library-Management-System/book"
	"github.com/mukul1234567/Library-Management-System/transaction"
	"github.com/mukul1234567/Library-Management-System/user"
)

const (
// versionHeader = "Accept"
)

func initRouter(dep dependencies) (router *mux.Router) {

	router = mux.NewRouter()
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)

	//User

	router.HandleFunc("/users", user.Create(dep.UserService)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.List(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.FindByID(dep.UserService)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.DeleteByID(dep.UserService)).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Update(dep.UserService)).Methods(http.MethodPut)

	//Book

	router.HandleFunc("/books", book.Create(dep.BookService)).Methods(http.MethodPost)
	router.HandleFunc("/books", book.List(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.FindByID(dep.BookService)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.DeleteByID(dep.BookService)).Methods(http.MethodDelete)
	router.HandleFunc("/books", book.Update(dep.BookService)).Methods(http.MethodPut)

	//Transaction

	router.HandleFunc("/book/issue", transaction.Create(dep.TransactionService)).Methods(http.MethodPost)
	router.HandleFunc("/book", transaction.List(dep.TransactionService)).Methods(http.MethodGet)
	router.HandleFunc("/book/return", transaction.Update(dep.TransactionService)).Methods(http.MethodPut)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
