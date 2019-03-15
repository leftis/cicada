package graphql

import (
	"context"
	"io/ioutil"

	"github.com/graph-gophers/graphql-go"
	"github.com/leftis/cicada/config"
	"github.com/leftis/cicada/models"
)

var Schema *graphql.Schema

type Resolver struct{}

func Init() {
	s, err := getSchema(config.App.CurrentDirectory + "/graphql/schema.graphql")
	if err != nil {
		panic(err)
	}
	Schema = graphql.MustParseSchema(s, &Resolver{}, graphql.UseStringDescriptions())
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

// Root resolvers
func (r *Resolver) GetAdministrator(
	ctx context.Context,
	args struct{ Id graphql.ID }) *models.AdministratorResolver {
	a := models.Administrator{}
	a.FirstBy(map[string]interface{}{"id": args.Id})
	return a.Resolver()
}
