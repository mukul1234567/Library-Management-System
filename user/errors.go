package user

import "errors"

var (
	errEmptyID      = errors.New("user ID must be present")
	errEmptyName    = errors.New("please fill all the details")
	errNoUsers      = errors.New("no Users present")
	errNoUserId     = errors.New("user is not present")
	errInvalidEmail = errors.New("invalid Email Entered")
	errInvalidMobNo = errors.New("invalid Mobile Number Entered")
	errInvalidAge   = errors.New("invalid Age Entered")
	errInvalidRole  = errors.New("invalid Role Entered")
)
