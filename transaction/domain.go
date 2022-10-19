package transaction

import "github.com/mukul1234567/Library-Management-System/db"

type UpdateRequest struct {
	ID         string `json:"id"`
	IssueDate  string `json:"issuedate"`
	DueDate    string `json:"duedate"`
	ReturnDate string `json:"returndate"`
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
}

type CreateRequest struct {
	ID         string `json:"id"`
	IssueDate  string `json:"issuedate"`
	DueDate    string `json:"duedate"`
	ReturnDate string `json:"returndate"`
	BookID     string `json:"book_id"`
	UserID     string `json:"user_id"`
}

type findByUserIDResponse struct {
	Transaction db.Transaction `json:"transaction"`
}

type findByBookIDResponse struct {
	Transaction db.Transaction `json:"transaction"`
}

type listResponse struct {
	Transaction []db.Transaction `json:"transactions"`
}

func (cr CreateRequest) Validate() (err error) {
	if cr.BookID == "" {
		return errEmptyBookID
	}
	if cr.UserID == "" {
		return errEmptyUserID
	}
	return
}

func (ur UpdateRequest) Validate() (err error) {
	if ur.BookID == "" {
		return errEmptyBookID
	}
	if ur.UserID == "" {
		return errEmptyUserID
	}
	return
}

//
