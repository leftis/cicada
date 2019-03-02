package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/leftis/cicada/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Administrator struct {
	Username 		string `db:"username" json:"username"`
	Password 		string `json:"password"`
	HashedPassword 	string `db:"hashed_password"`
	jwt.StandardClaims
}

func (admin *Administrator) Authenticate() *Administrator {
	err := db.SQLX.Get(admin, "SELECT * FROM administrators WHERE username = $1", admin.Username)
	if err != nil {
		log.Fatal(err)
	}

	if admin.hashedPasswordMatch() {
		return admin
	} else {
		return nil
	}
}

func (admin *Administrator) hashedPasswordMatch() bool {
	err := bcrypt.CompareHashAndPassword([]byte(admin.HashedPassword), []byte(admin.Password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
