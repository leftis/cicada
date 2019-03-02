package graphql

import (
	"context"
	"github.com/graph-gophers/graphql-go"
	"github.com/leftis/cicada/db"
	"github.com/leftis/cicada/models"
)

var Schema *graphql.Schema

const schema = `
	schema {
		query: Query
	}
	
	type Query {
		administrator(id: ID!): Administrator!
	}
	
	type Administrator {
		id: ID!
		username: String!
	}
`

type Resolver struct{}

type administratorResolver struct {
	Administrator models.Administrator
}

func (r administratorResolver) Id() graphql.ID {
	return graphql.ID(r.Administrator.StandardClaims.Id)
}

func (r administratorResolver) Username() string {
	return r.Administrator.Username
}

func (r *Resolver) Administrator(ctx context.Context, args struct { ID graphql.ID }) *administratorResolver {
	admin := models.Administrator{}
	//admin := models.Administrator.findBy()
	db.SQLX.Get(&admin, "SELECT * FROM administrators WHERE id = $1", args.ID)
	return &administratorResolver{admin}
}

func (r *Resolver) Administrators() *[]*administratorResolver {
	rows, _ := db.SQLX.Query("SELECT * FROM administrators")
	resolvers := make([]*administratorResolver, 0)
	for rows.Next() {
		a := models.Administrator{}
		rows.Scan(&a)
		resolvers = append(resolvers, &administratorResolver{a})
	}
	return &resolvers
}

func Init() {
	Schema = graphql.MustParseSchema(schema, &Resolver{})
}