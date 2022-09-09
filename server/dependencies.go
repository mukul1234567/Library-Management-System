package server

import _ "github.com/mukul1234567/Library-Management-System/app"

type dependencies struct {
	// 	UserService user.Service
}

func initDependencies() (dependencies, error) {
	// 	appDB := app.GetDB()
	// 	logger := app.GetLogger()
	// 	// dbStore := db.NewStorer(appDB)

	// 	// userService := user.NewService(dbStore, logger)

	return dependencies{
		// UserService: userService,
	}, nil

}
