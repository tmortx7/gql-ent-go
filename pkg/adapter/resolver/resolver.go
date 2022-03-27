package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/tmortx7/gql-ent-go/ent"
	"github.com/tmortx7/gql-ent-go/graph/generated"
	"github.com/tmortx7/gql-ent-go/pkg/adapter/controller"
	"github.com/tmortx7/gql-ent-go/pkg/adapter/directives"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is a context struct
type Resolver struct {
	client     *ent.Client
	controller controller.Controller
}

// NewSchema creates NewExecutableSchema
func NewSchema(client *ent.Client, controller controller.Controller) graphql.ExecutableSchema {
	c := generated.Config{
		Resolvers: &Resolver{
			client:     client,
			controller: controller,
		},
	}
	c.Directives.Binding = directives.Binding

	return generated.NewExecutableSchema(c)
}
