package db

import "errors"

var (
	ErrUserNotExist = errors.New("Users does not exist in db")
)
var (
	ErrBookNotExist = errors.New("Books does not exist in db")
)
