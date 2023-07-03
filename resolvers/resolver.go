package resolvers

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/kameshsampath/fruitsgql"
	"github.com/kameshsampath/fruitsgql/ent"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return fruitsgql.NewExecutableSchema(fruitsgql.Config{
		Resolvers: &Resolver{client: client},
	})
}
