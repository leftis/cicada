package main

import (
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/db"
	"os"
)

func main() {
	config.Init()
	db.Init()
	defer db.Close()

	run(os.Args)
	return
}
