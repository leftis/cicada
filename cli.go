package main

import (
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/db"
	"github.com/leftis/cicada/server"
	"gopkg.in/urfave/cli.v1"
	"log"
)

func runApplication(_ *cli.Context) error {
	server.Init()
	return nil
}

func runMigrations(_ *cli.Context) error {
	db.DB.Migrate(config.App.CurrentDirectory)
	return nil
}

func run(args cli.Args) {
	app := cli.NewApp()

	app.Name = "cicada"
	app.Usage = "Cicada - the best tool"

	app.Commands = []cli.Command{
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "run our application",
			Action:  runApplication,
		},
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Usage:   "run our migrations",
			Action:  runMigrations,
		},
	}

	err := app.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}
