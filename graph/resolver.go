package graph

import (
	"github.com/BradHacker/chungus/ent"
	"github.com/BradHacker/chungus/graph/generated"

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
