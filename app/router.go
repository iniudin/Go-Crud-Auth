package app

import (
	"go-crud-auth/controller"

	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes(server *Server) {
	authController := controller.NewAuthController()
	server.Echo.Use(middleware.Logger())

	server.Echo.GET("/", authController.Login)

}
