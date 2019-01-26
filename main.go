package main

import (
	"github.com/leftis/cicada/configuration"
	"github.com/leftis/cicada/server"
)

func main() {
	app := configuration.Init()
	//database.Init(app)
	server.Init(app)
	return
}
