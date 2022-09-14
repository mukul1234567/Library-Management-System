package transaction

import (
	"context"
	"fmt"

	// "github.com/mukul1234567/Library-Management-System/book"
	"github.com/google/uuid"
	"github.com/mukul1234567/Library-Management-System/db"
	"go.uber.org/zap"
)

// var Availablecopiesval int

type Service interface {
	list(ctx context.Context) (response listResponse, err error)
	create(ctx context.Context, req createRequest) (err error)
	// findByID(ctx context.Context, id string) (response findByIDResponse, err error)
	// deleteByID(ctx context.Context, id string) (err error)
	update(ctx context.Context, req updateRequest) (err error)
}

type transactionService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

func (cs *transactionService) list(ctx context.Context) (response listResponse, err error) {
	fmt.Println()
	transactions, err := cs.store.ListTransactions(ctx)
	if err == db.ErrTransactionNotExist {
		cs.logger.Error("No transaction present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing transactions", "err", err.Error())
		return
	}

	response.Transaction = transactions
	return
}

func (cs *transactionService) create(ctx context.Context, c createRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Errorw("Invalid request for transaction create", "msg", err.Error(), "transaction", c)
		return
	}
	uuidgen := uuid.New()
	c.ID = uuidgen.String()
	err = cs.store.CreateTransaction(ctx, &db.Transaction{

		ID:         c.ID,
		IssueDate:  c.IssueDate,
		ReturnDate: c.ReturnDate,
		DueDate:    c.DueDate,
		BookID:     c.BookID,
		UserID:     c.UserID,
	})
	if err != nil {
		cs.logger.Error("Error creating transaction", "err", err.Error())
		return
	}
	// Availablecopiesval = Availablecopiesval - 1
	return
}

func (cs *transactionService) update(ctx context.Context, c updateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateTransaction(ctx, &db.Transaction{
		ID:         c.ID,
		IssueDate:  c.IssueDate,
		ReturnDate: c.ReturnDate,
		DueDate:    c.DueDate,
		BookID:     c.BookID,
		UserID:     c.UserID,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

// func (cs *userService) findByID(ctx context.Context, id string) (response findByIDResponse, err error) {
// 	user, err := cs.store.FindUserByID(ctx, id)
// 	if err == db.ErrUserNotExist {
// 		cs.logger.Error("No user present", "err", err.Error())
// 		return response, errNoUserId
// 	}
// 	if err != nil {
// 		cs.logger.Error("Error finding user", "err", err.Error(), "id", id)
// 		return
// 	}

// 	response.User = user
// 	return
// }

// func (cs *userService) deleteByID(ctx context.Context, id string) (err error) {
// 	err = cs.store.DeleteUserByID(ctx, id)
// 	if err == db.ErrUserNotExist {
// 		cs.logger.Error("user Not present", "err", err.Error(), "id", id)
// 		return errNoUserId
// 	}
// 	if err != nil {
// 		cs.logger.Error("Error deleting user", "err", err.Error(), "id", id)
// 		return
// 	}

// 	return
// }

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &transactionService{
		store:  s,
		logger: l,
	}
}
