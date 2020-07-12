package main

import (
	"log"

	"github.com/raygervais/athenaeum/cmd/api/router"
	"github.com/raygervais/athenaeum/cmd/models"
	"github.com/raygervais/athenaeum/pkg/application"
	"github.com/raygervais/athenaeum/pkg/logger"
	"github.com/raygervais/athenaeum/pkg/server"
)

func main() {
	app, err := application.Get()
	app.Database.InitializeTables(&models.Book{})

	if err != nil {
		log.Fatal(err.Error())
	}

	server := server.Get().ConnectRouter(router.SetupRouter(app))
	logger.Info.Printf("starting server at %s", app.Config.GetAPIPort())
	if err := server.Start(app.Config.GetAPIPort()); err != nil {
		logger.Error.Fatal(err.Error())
	}

}
