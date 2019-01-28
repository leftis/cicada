package models

type Administrator struct {
	Id       string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}
