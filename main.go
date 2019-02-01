package main

import (
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/db"
	"os"
)

func main() {
	config.Init()
	db.Init()
	defer db.DB.Conn.Close()

	//TODO: custom Close function on connection.go to handle Close

	run(os.Args)
	return
}
