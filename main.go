package main

import (
	"github.com/leftis/cicada/configuration"
	"github.com/leftis/cicada/db"
	"os"
)

var (
	appConfig configuration.App
	database  *db.Connection
)

func main() {
	// Grab configuration
	appConfig = configuration.Init()

	// Establish database connection
	database = db.Init(appConfig)
	defer database.Conn.Close()

	// Run command line interface
	run(os.Args)

	return
}
