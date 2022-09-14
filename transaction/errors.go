package transaction

import "errors"

var (
	errEmptyUserID = errors.New("User ID must be present")
	errEmptyBookID = errors.New("Book ID must be present")
	errNoUsers     = errors.New("No Users present")
	errNoUserId    = errors.New("User is not present")
	errNoBookId    = errors.New("Book is not present")
)
//