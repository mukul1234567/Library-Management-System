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
	router.HandleFunc("/login", user.Login(dep.UserService)).Methods(http.MethodPost)

	router.HandleFunc("/users", user.Authorize(user.Create(dep.UserService), 1)).Methods(http.MethodPost)
	router.HandleFunc("/users", user.Authorize(user.List(dep.UserService), 1)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.Authorize(user.FindByID(dep.UserService), 2)).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", user.Authorize(user.DeleteByID(dep.UserService), 1)).Methods(http.MethodDelete)
	router.HandleFunc("/users", user.Authorize(user.Update(dep.UserService), 2)).Methods(http.MethodPut)

	//Book

	router.HandleFunc("/books", user.Authorize(book.Create(dep.BookService), 1)).Methods(http.MethodPost)
	router.HandleFunc("/books", user.Authorize(book.List(dep.BookService), 2)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", user.Authorize(book.FindByID(dep.BookService), 2)).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", user.Authorize(book.DeleteByID(dep.BookService), 1)).Methods(http.MethodDelete)
	router.HandleFunc("/books", user.Authorize(book.Update(dep.BookService), 1)).Methods(http.MethodPut)

	//Transaction

	router.HandleFunc("/book/issue", user.Authorize(transaction.Create(dep.TransactionService), 1)).Methods(http.MethodPost)
	router.HandleFunc("/book", user.Authorize(transaction.List(dep.TransactionService), 1)).Methods(http.MethodGet)
	router.HandleFunc("/book/return", user.Authorize(transaction.Update(dep.TransactionService), 1)).Methods(http.MethodPut)
	router.HandleFunc("/book/{book_id}", user.Authorize(transaction.FindByBookID(dep.TransactionService), 1)).Methods(http.MethodGet)

	return
}

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	api.Success(rw, http.StatusOK, api.Response{Message: "pong"})
}
