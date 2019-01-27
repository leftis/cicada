package database

import (
	"database/sql"
	"fmt"
	"github.com/leftis/cicada/configuration"
	_ "github.com/lib/pq"
	"log"
)

type connection struct {
	url string

	conn *sql.DB
	conf *configuration.Database
}

var (
	Database connection
)

func Init(app configuration.App) {
	assignDatabaseConfig(&app.Config.Database)
	createDatabaseUlr()
	createConnection()
	//checkMigrations()
}

func assignDatabaseConfig(dbConf *configuration.Database) {
	Database.conf = dbConf
}

func createConnection() {
	fmt.Println("Connecting to " + Database.url)
	conn, err := sql.Open(Database.conf.Driver,  Database.url)

	//defer conn.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	Database.conn = conn
}

func createDatabaseUlr() {
	c := Database.conf
	Database.url = "%s://%s:%s@%s:%s/%s?sslmode=%s"
	Database.url = fmt.Sprintf(Database.url, c.Driver, c.User, c.Pass, c.Host, c.Port, c.Name, c.SSL)
}