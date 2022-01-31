package graph

import (
	"github.com/BradHacker/Br4vo6ix/ent"
	"github.com/BradHacker/Br4vo6ix/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	client *ent.Client
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	GQLConfig := generated.Config{
		Resolvers: &Resolver{
			client: client,
		},
	}
	return generated.NewExecutableSchema(GQLConfig)
}
