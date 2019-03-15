package main

import (
	"os"

	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/db"
)

func main() {
	config.Init()
	db.Init()
	defer db.Close()

	run(os.Args)
	return
}
