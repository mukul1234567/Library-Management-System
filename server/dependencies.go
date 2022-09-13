package server

import (
	"github.com/mukul1234567/Library-Management-System/app"
	"github.com/mukul1234567/Library-Management-System/book"
	"github.com/mukul1234567/Library-Management-System/db"
	"github.com/mukul1234567/Library-Management-System/user"
)

type dependencies struct {
	UserService user.Service
	BookService book.Service
}

func initDependencies() (dependencies, error) {
	appDB := app.GetDB()
	logger := app.GetLogger()
	dbStore := db.NewStorer(appDB)

	userService := user.NewService(dbStore, logger)
	bookService := book.NewService(dbStore, logger)

	return dependencies{
		UserService: userService,
		BookService: bookService,
	}, nil

}

// {
//     "id": "1",
//     "name": "Harry Potter",
//     "author": "Charles Dickens",
//     "price": 670,
//     "totalcopies": 50,
//     "status": "Available",
//     "availablecopies": 20
// }

// {
//     "first_name": "Sachin",
//     "last_name": "Tendulkar",
//     "gender": "Male",
//     "age": 47,
//     "address": "Pune",
//     "email": "sachin@gmail.com",
//     "password": "sachin@123",
//     "mob_no": "9954861237",
//     "role":"user"
// }
