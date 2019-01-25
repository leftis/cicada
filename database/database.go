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

var Database = &db{}

func Init() {
	createConnection()
	//checkMigrations()
}

func createConnection() {
	databaseUrl := dbURL()

	fmt.Println("Connecting to "+databaseUrl)
	conn, err := sql.Open(configuration.Config()["database"]["driver"], databaseUrl)

	//defer conn.Close()

	if err != nil {
		log.Fatal("could not get a connection")
	}

	if err := conn.Ping(); err != nil {
		log.Fatal("could not establish a good connection")
	}

	Database.conn = conn

}

func dbURL() string {
	c := configuration.Config()["database"]
	connStr := "%s://%s:%s@%s:%s/%s?sslmode=%s"
	connStr = fmt.Sprintf(connStr, c["driver"], c["user"], c["pass"], c["host"], c["port"], c["name"], c["ssl"])
	return connStr
}