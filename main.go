package main

import (
	"Secreto/app"
	"Secreto/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	app.NewDatabase()

	authController := controller.NewAuthController()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Router -> Handler

	e.GET("/login", authController.Login)

	e.Logger.Fatal(e.Start(":3000"))
}
