package main

import (
	"go-crud-auth/app"
	"go-crud-auth/config"
	"go-crud-auth/helper"
	"go-crud-auth/router"
)

func main() {
	config, err := config.NewConfig()
	helper.PanicError(err)

	server := app.NewServer(config)
	router.NewRoutes(server)
	server.Start(config.Port)
}
