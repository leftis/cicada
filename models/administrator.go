package models

import (
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/graph-gophers/graphql-go"
	h "github.com/leftis/cicada/helpers"

	"github.com/dgrijalva/jwt-go"
	"github.com/leftis/cicada/db"
	"golang.org/x/crypto/bcrypt"
)

type Administrator struct {
	Id             uint64 `db:"id"`
	Username       string `db:"username" json:"username"`
	Password       string `json:"password"`
	HashedPassword string `db:"hashed_password"`
	jwt.StandardClaims
}

func ConditionedSelectQuery(columns string, table string, conditions map[string]interface{}) sq.SelectBuilder {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return psql.Select("*").From(table).Where(conditions)
}

func (m *Administrator) FindBy(conditions map[string]interface{}) error {
	sql, args, err := ConditionedSelectQuery("*", "administrators", conditions).ToSql()
	if err != nil {
		log.Fatal(err)
	}
	return db.SQLX.Get(m, sql, args...)
}

func (m *Administrator) FirstBy(conditions map[string]interface{}) error {
	sql, args, err := ConditionedSelectQuery("*", "administrators", conditions).Limit(1).ToSql()
	if err != nil {
		log.Fatal(err)
	}
	return db.SQLX.Get(m, sql, args...)
}

func (m *Administrator) Authenticate() *Administrator {
	conditions := map[string]interface{}{"username": m.Username}
	err := m.FirstBy(conditions)

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
	M *Administrator
}

func (m *Administrator) Resolver() *AdministratorResolver {
	return &AdministratorResolver{m}
}

func (u *AdministratorResolver) Id() *graphql.ID {
	return h.GqlIDP(u.M.Id)
}

func (u *AdministratorResolver) Username() string {
	return u.M.Username
}
