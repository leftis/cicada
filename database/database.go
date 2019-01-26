package database

import (
	"database/sql"
	"fmt"
	"github.com/leftis/cicada/configuration"
	_ "github.com/lib/pq"
	"log"
)

type db struct {
	conn *sql.DB
}

var (
	Database db
	config map[string]string
	databaseUrl string
)

func Init(app configuration.App) {
	assignDatabaseConfig(app)
	createDatabaseUlr()
	createConnection()
	//checkMigrations()
}

func assignDatabaseConfig(app configuration.App) {
	config = app.Config["database"]
}

func createConnection() {
	fmt.Println("Connecting to " + databaseUrl)
	conn, err := sql.Open(config["driver"], databaseUrl)

	//defer conn.Close()

	if err != nil {
		log.Fatal("could not get a connection")
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("could not establish a good connection")
	}

	Database.conn = conn

}

func createDatabaseUlr() {
	c := config
	databaseUrl := "%s://%s:%s@%s:%s/%s?sslmode=%s"
	databaseUrl = fmt.Sprintf(databaseUrl, c["driver"], c["user"], c["pass"], c["host"], c["port"], c["name"], c["ssl"])
}