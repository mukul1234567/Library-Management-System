package db

import (
	"context"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gender    string `db:"gender"`
	Address   string `db:"address"`
	Age       int    `db:"age"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	MobileNum string `db:"mob_no"`
	Role      string `db:"role"`
	deleted   int
}

type Userlist struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Gender    string `db:"gender"`
	Address   string `db:"address"`
	Age       int    `db:"age"`
	Email     string `db:"email"`
	MobileNum string `db:"mob_no"`
	Role      string `db:"role"`
}

const (
	createUserQuery = `INSERT INTO users (id,first_name, last_name, gender,age,address,email,password,mob_no,role)
	VALUES(?, ?,?,?,?,?,?,?,?,?)`
	listUsersQuery          = `SELECT id,first_name, last_name, gender,age,address,email,password,mob_no,role FROM users`
	showUsersQuery          = `SELECT id,first_name, last_name, gender,age,address,email,mob_no,role FROM users`
	findUserByIDQuery       = `SELECT first_name, last_name, gender,age,address,email,mob_no,role FROM users WHERE id = ?`
	deleteUserByIDQuery     = `DELETE FROM users WHERE id = ?`
	updateUserQuery         = `UPDATE users SET first_name=?, last_name=?, gender=?, age=?, address=?, password=?, mob_no=? where id = ?`
	updateUserPasswordQuery = `UPDATE users SET password = ? WHERE id =?`
)

func (s *store) CreateUser(ctx context.Context, user *User) (err error) {

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			createUserQuery,
			user.ID,
			user.FirstName,
			user.LastName,
			user.Gender,
			user.Age,
			user.Address,
			user.Email,
			HashPassword(user.Password),
			user.MobileNum,
			user.Role,
		)
		return err
	})
}

func (s *store) ListUsers(ctx context.Context) (users []User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &users, listUsersQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}
	return
}

func (s *store) ShowUsers(ctx context.Context) (users []Userlist, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.SelectContext(ctx, &users, showUsersQuery)
	})
	if err == sql.ErrNoRows {
		return users, ErrUserNotExist
	}
	return
}

func (s *store) FindUserByID(ctx context.Context, id string) (user User, err error) {
	err = WithDefaultTimeout(ctx, func(ctx context.Context) error {
		return s.db.GetContext(ctx, &user, findUserByIDQuery, id)
	})
	if err == sql.ErrNoRows {
		return user, ErrUserNotExist
	}
	return
}

func (s *store) DeleteUserByID(ctx context.Context, id string) (err error) {
	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		res, err := s.db.Exec(deleteUserByIDQuery, id)
		if err != nil {
			return err
		}
		cnt, err := res.RowsAffected()
		if cnt == 0 {
			return ErrUserNotExist
		}
		if err != nil {
			return err
		}
		return err
	})
}

func (s *store) UpdateUser(ctx context.Context, user *User) (err error) {
	// now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserQuery,
			user.FirstName,
			user.LastName,
			user.Gender,
			user.Age,
			user.Address,
			user.Password,
			user.MobileNum,
			user.ID,
		)
		return err
	})
}

func (s *store) UpdatePassword(ctx context.Context, user *User) (err error) {
	// now := time.Now()

	return Transact(ctx, s.db, &sql.TxOptions{}, func(ctx context.Context) error {
		_, err = s.db.Exec(
			updateUserPasswordQuery,
			user.Password,
			user.ID,
		)
		return err
	})
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	fmt.Println(err)
	return string(bytes)
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
