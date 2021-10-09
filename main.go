package main

import (
	"go-crud-auth/app"
	"go-crud-auth/config"
	"go-crud-auth/helper"
)

func main() {
	config, err := config.NewConfig()
	helper.PanicError(err)

	app := app.NewServer(config)
	app.Start(config.Port)
}
