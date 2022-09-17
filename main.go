package main

import (
	"os"

	"github.com/urfave/cli"

	"github.com/mukul1234567/Library-Management-System/app"
	"github.com/mukul1234567/Library-Management-System/config"
	"github.com/mukul1234567/Library-Management-System/db"
	"github.com/mukul1234567/Library-Management-System/server"
)

func main() {
	config.Load()
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "Golang Library app"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				err := db.RunMigrations()
				return err
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
