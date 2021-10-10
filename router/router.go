package router

import (
	"go-crud-auth/app"
	"go-crud-auth/controller"

	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes(server *app.Server) {
	authController := controller.NewAuthController(server)
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/", authController.Login)

}
