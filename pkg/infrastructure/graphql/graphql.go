package graphql

import (
	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/tmortx7/gql-ent-go/ent"
	"github.com/tmortx7/gql-ent-go/pkg/adapter/controller"
	"github.com/tmortx7/gql-ent-go/pkg/adapter/resolver"
)

// NewServer generates graphql server
func NewServer(client *ent.Client, controller controller.Controller) *handler.Server {
	srv := handler.NewDefaultServer(resolver.NewSchema(client, controller))
	srv.Use(entgql.Transactioner{TxOpener: client})

	return srv
}
