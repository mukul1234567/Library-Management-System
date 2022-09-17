package db

import (
	"context"
	"database/sql"
	// "fmt"

	// "fmt"
	"time"
	// "github.com/mukul1234567/Library-Management-System/db"
)

type Transaction struct {
	ID         string `db:"id"`
	IssueDate  int    `db:"issuedate"`
	ReturnDate int    `db:"returndate"`
	DueDate    int    `db:"duedate"`
	BookID     string `db:"book_id"`
	UserID     string `db:"user_id"`
}

var Availablecopiesval int

const (
	createTransactionQuery = `INSERT INTO transactions (id,issuedate,duedate,returndate,book_id,user_id)
	VALUES(?, ?,?,?,?,?)`
	checkTransactionQuery        = `SELECT COUNT(*) FROM transactions WHERE book_id = ? AND user_id =? AND returndate=0`
	listTransactionsQuery        = `SELECT * FROM transactions`
	findTransactionByBookIDQuery = `SELECT * FROM transactions WHERE book_id = ?`
	findTransactionByUserIDQuery = `SELECT FROM transactions WHERE user_id = ?`
	updateTransactionQuery       = `UPDATE transactions SET returndate=? WHERE book_id = ? AND user_id =?`
	issueCopiesQuery             = `UPDATE books SET availablecopies=availablecopies-1 WHERE id = ? AND availablecopies>0`
	returnCopiesQuery            = `UPDATE books SET availablecopies=availablecopies+1 WHERE id = ?`
)

func (s *store) CreateTransaction(ctx context.Context, transaction *Transaction) (err error) {
	now := time.Now().UTC().Unix()
	transaction.DueDate = int(now) + 864000
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		// res, err := s.db.Exec(
		// 	checkTransactionQuery,
		// 	transaction.BookID,
		// 	transaction.UserID,
		// )
		// fmt.Println(res)
		// if int(res) == 0 {
		// 	return ErrBookAlreadyIssued
		// }

		if err != nil {
			return err
		}
		_, err = s.db.Exec(
			createTransactionQuery,
			transaction.ID,
			now,
			transaction.DueDate,
			0,
			transaction.BookID,
			transaction.UserID,
		)
		if err != nil {
			return err
		}
		_, err := s.db.Exec(
			issueCopiesQuery,
			transaction.BookID,
		)
		// cnt,err:=res.RowsAffected()
		// if cnt!=0{
		// 	api.Success(rw, http.StatusCreated, api.Response{Message: "Created Successfully"})
		// }
		return err
	})
}

func (s *store) ListTransactions(ctx context.Context) (transactions []Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &transactions, listTransactionsQuery)
	})
	if err == sql.ErrNoRows {
		return transactions, ErrTransactionNotExist
	}
	return
}

func (s *store) UpdateTransaction(ctx context.Context, transaction *Transaction) (err error) {
	now := time.Now().UTC().Unix()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateTransactionQuery,
			now,
			transaction.BookID,
			transaction.UserID,
		)
		if err != nil {
			return err
		}
		_, err = s.db.Exec(
			returnCopiesQuery,
			transaction.BookID,
		)
		return err
	})
}

func (s *store) FindTransactionByBookID(ctx context.Context, bookid string) (transaction Transaction, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &transaction, findTransactionByBookIDQuery, bookid)
	})
	if err == sql.ErrNoRows {
		return transaction, ErrTransactionNotExist
	}
	return
}

// CREATE TABLE Users
// (

//   `id` VARCHAR(10) NOT NULL,
//   `first_name` VARCHAR(20) NOT NULL,
//   `last_name` VARCHAR(20) NOT NULL,
//   `gender` VARCHAR(10) NOT NULL,
//   `age` INT NOT NULL,
//   `address` VARCHAR(50) NOT NULL,
//   `email` VARCHAR(30) NOT NULL,
//   `password` VARCHAR(20) NOT NULL,
//   `mob_no` VARCHAR(50) NOT NULL,
//   `role` VARCHAR(10) NOT NULL,
//   PRIMARY KEY (id),
//   UNIQUE (email)
// );

// IF EXISTS (SELECT * FROM transactions WHERE book_id = 911bb240-4981-4e2c-b2f4-d1f3c7aa3268 AND user_id = d563e110-ac05-4904-be9c-1cbf42939833 AND returndate IS NOT NULL)
// BEGIN
// INSERT INTO transactions (book_id,user_id)
// 	VALUES(911bb240-4981-4e2c-b2f4-d1f3c7aa3268,d563e110-ac05-4904-be9c-1cbf42939833)
// END

// IF EXISTS
// (
//     SELECT * FROM transactions WHERE book_id = "911bb240-4981-4e2c-b2f4-d1f3c7aa3268" AND user_id = "d563e110-ac05-4904-be9c-1cbf42939833" AND returndate IS NOT NULL
// )
//     BEGIN
// 	INSERT INTO transactions (book_id,user_id) VALUES("911bb240-4981-4e2c-b2f4-d1f3c7aa3268","d563e110-ac05-4904-be9c-1cbf42939833")
// END;

// INSERT INTO transactions (book_id,user_id) VALUES("911bb240-4981-4e2c-b2f4-d1f3c7aa3268","d563e110-ac05-4904-be9c-1cbf42939833")
// SELECT * FROM transactions
// WHERE book_id = "911bb240-4981-4e2c-b2f4-d1f3c7aa3268" AND user_id = "d563e110-ac05-4904-be9c-1cbf42939833" AND returndate IS NOT NULL;

// {
//     "first_name": "Perooop",
//     "last_name": "Ghhhhgfff",
//     "gender": "Male",
//     "age": 47,
//     "address": "Pune",
//     "email": "cntst@gmail.com",
//     "password": "dtest@123",
//     "mob_no": "9865327485",
//     "role":"user"
// }
