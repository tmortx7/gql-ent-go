package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/tmortx7/gql-ent-go/config"
	"github.com/tmortx7/gql-ent-go/ent"
	"github.com/tmortx7/gql-ent-go/pkg/adapter/controller"
	"github.com/tmortx7/gql-ent-go/pkg/infrastructure/datastore"
	"github.com/tmortx7/gql-ent-go/pkg/infrastructure/graphql"
	"github.com/tmortx7/gql-ent-go/pkg/infrastructure/router"
	"github.com/tmortx7/gql-ent-go/pkg/registry"
)

func main() {
	config.ReadConfig(config.ReadConfigOption{})

	client := newDBClient()
	ctrl := newController(client)

	srv := graphql.NewServer(client, ctrl)
	e := router.New(srv)

	e.Logger.Fatal(e.Start(":" + config.C.Server.Port))
}

func newDBClient() *ent.Client {
	client, err := datastore.NewClient()
	if err != nil {
		log.Fatalf("failed opening postgres client: %v", err)
	}

	return client
}

func newController(client *ent.Client) controller.Controller {
	r := registry.New(client)
	return r.NewController()
}