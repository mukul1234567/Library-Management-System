package db

import "errors"

var (
	ErrUserNotExist = errors.New("Users does not exist in db")
)
var (
	ErrBookNotExist = errors.New("Books does not exist in db")
)

var (
	ErrTransactionNotExist = errors.New("Transactions does not exist in db")
)
var (
	ErrBookNotAvailable = errors.New("Book is not available")
)
var(
	ErrBookAlreadyIssued = errors.New("Book has already been issued")
)
