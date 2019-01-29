package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/leftis/cicada/configuration"
	_ "github.com/lib/pq"
	"log"
)

type Connection struct {
	Conn *sql.DB
	Conf *configuration.Database
}

var (
	Database Connection
)

func Init(app configuration.App) *Connection {
	Database.assignConfiguration(&app.Config.Database)
	Database.open()

	return &Database
}

func (c *Connection) url() string {
	d := c.Conf
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", d.Driver, d.User, d.Pass, d.Host, d.Port, d.Name, d.SSL)
}

func (c *Connection) assignConfiguration(dbConf *configuration.Database) {
	c.Conf = dbConf
}

func (c *Connection) open() {

	conn, err := sql.Open(c.Conf.Driver, c.url())

	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	c.Conn = conn
}

type migrationsLog struct{}

func(m migrationsLog) Printf(format string, v ...interface {}) {
	fmt.Printf(format, v)
}

func (m migrationsLog) Verbose() bool {
	return true
}

func (c *Connection) Migrate(currentDirectory string) {
	mlog := migrationsLog{}
	driver, err := postgres.WithInstance(c.Conn, &postgres.Config{})

	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/db/migrations", currentDirectory), c.Conf.Name, driver)

	m.Log = mlog

	if err != nil {
		log.Fatal(err)
	}

	m.Steps(2)
}
