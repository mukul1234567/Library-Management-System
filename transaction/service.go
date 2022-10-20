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
	create(ctx context.Context, req CreateRequest) (err error)
	findByBookID(ctx context.Context, id string) (response findByBookIDResponse, err error)
	findByUserID(ctx context.Context, id string) (response findByUserIDResponse, err error)
	// deleteByID(ctx context.Context, id string) (err error)
	update(ctx context.Context, req UpdateRequest) (err error)
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

func (cs *transactionService) create(ctx context.Context, c CreateRequest) (err error) {
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

	return
}

func (cs *transactionService) update(ctx context.Context, c UpdateRequest) (err error) {
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

func (cs *transactionService) findByBookID(ctx context.Context, id string) (response findByBookIDResponse, err error) {
	transaction, err := cs.store.FindTransactionByBookID(ctx, id)
	if err == db.ErrTransactionNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "id", id)
		return
	}

	response.Transaction = transaction
	return
}

func (cs *transactionService) findByUserID(ctx context.Context, id string) (response findByUserIDResponse, err error) {
	transaction, err := cs.store.FindTransactionByUserID(ctx, id)
	if err == db.ErrTransactionNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "id", id)
		return
	}

	response.Transaction = transaction
	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &transactionService{
		store:  s,
		logger: l,
	}
}
