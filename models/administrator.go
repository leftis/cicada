package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/graph-gophers/graphql-go"
	"github.com/leftis/cicada/db"
	"github.com/leftis/cicada/helpers"
	"golang.org/x/crypto/bcrypt"
	sq "github.com/Masterminds/squirrel"
	"log"
)

type Administrator struct {
	ID              uint   `db:"id"`
	Username 		string `db:"username" json:"username"`
	Password 		string `json:"password"`
	HashedPassword 	string `db:"hashed_password"`
	jwt.StandardClaims
}

func (m *Administrator) GetBy(pair helpers.Pair) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	q := psql.Select("*").From("administrators").Where().Limit(1)
	sql, args, _ := q.ToSql()
	println(args)
	println(sql)

	return db.SQLX.Get(m, sql, pair.Key.(string), pair.Value.(string))
}

func (m *Administrator) Authenticate() *Administrator {
	err := m.GetBy(helpers.Pair{"username",m.Username})

	if err != nil {
		log.Fatal(err)
	}

	if m.hashedPasswordMatch() {
		return m
	} else {
		return nil
	}
}

func (m *Administrator) hashedPasswordMatch() bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.HashedPassword), []byte(m.Password))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// Graphql Specifics
type AdministratorResolver struct {
	M Administrator
}

func (u *AdministratorResolver) Id() *graphql.ID {
	return helpers.GqlIDP(u.M.ID)
}

func (u *AdministratorResolver) Username() string {
	return u.M.Username
}
