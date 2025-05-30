package main

import (
	"os"

	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/application/grpc"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/config"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/core"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/infrastructure/quickwit"
	"github.com/Courtcircuits/tad-beep/pocs/inter-service/messages/infrastructure/sql"
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

	// infrastructures
	search_engine, err := quickwit.NewQuickwit(conf.Quickwit)
	if err != nil {
		optique.Error(err.Error())
		cycle.Stop()
		os.Exit(1)
	}
	db, err := sql.NewSql(conf.SQL)
	if err != nil {
		optique.Error(err.Error())
		cycle.Stop()
		os.Exit(1)
	}
	// services
	channel_service := core.NewChannelService()
	message_service := core.NewMessageService(channel_service, db, search_engine)
	// application
	grpc_server := grpc.NewGRPCServer(conf.GRPC, message_service, channel_service)


	cycle.AddApplication(grpc_server)
	cycle.AddRepository(search_engine)
	cycle.AddRepository(db)

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
