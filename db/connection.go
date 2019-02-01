package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/leftis/cicada/config"
	_ "github.com/lib/pq"
	"log"
)

var (
	DB Connection
	SQLX *sqlx.DB
)

type Connection struct {
	Conn *sql.DB
	Conf *config.Database
}

type migrationsLog struct{}

func Init() {
	DB.Conf = &config.App.Config.Database
	DB.open()
	SQLX = sqlx.NewDb(DB.Conn, DB.Conf.Driver)
}

func (c *Connection) url() string {
	d := c.Conf
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s", d.Driver, d.User, d.Pass, d.Host, d.Port, d.Name, d.SSL)
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

func(m migrationsLog) Printf(format string, v ...interface {}) {
	fmt.Printf(format, v)
}

func (m migrationsLog) Verbose() bool {
	return true
}