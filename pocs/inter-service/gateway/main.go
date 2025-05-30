package main

import (
	"os"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/application/graphql"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/config"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/core"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/gateway/infrastructure/grpc"
	"github.com/optique-dev/optique"
)

// @title Optique application TO CHANGE
// @version 1.0
// @description This is a sample application
// @contact.name Courtcircuits
// @contact.url https://github.com/Courtcircuits
// @contact.email tristan-mihai.radulescu@etu.umontpellier.fr
func main() {
	conf, err := config.LoadConfig()

	if err != nil {
		config.HandleError(err)
	}
	cycle := NewCycle()

	//infrastructure
	grpc_client, err := grpc.NewGRPCClient(conf.GRPC)
	if err != nil {
		optique.Error(err.Error())
		cycle.Stop()
		os.Exit(1)
	}

	// core
	message_service := core.NewMessageService(grpc_client)

	// application
	http_server := graphql.NewHttp(conf.GraphQL)
	graphql_handler := graphql.NewGraphQL(message_service)
	http_server.WithHandler(graphql_handler)

	cycle.AddApplication(http_server)
	cycle.AddRepository(grpc_client)

	if conf.Bootstrap {
		err := cycle.Setup()
		if err != nil {
			optique.Error(err.Error())
			cycle.Stop()
			os.Exit(1)
		}
	}

	err = cycle.Ignite()
}
