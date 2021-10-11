package router

import (
	"go-crud-auth/app"
	"go-crud-auth/controller"
	"go-crud-auth/repository"
	"go-crud-auth/service"

	"github.com/labstack/echo/v4/middleware"
)

func NewRoutes(server *app.Server) {

	// User
	userRepository := repository.NewUserRepository(server)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	server.Echo.Use(middleware.Logger())

	u := server.Echo.Group("/users")
	u.POST("/", userController.Create)
	u.PUT("/:userId", userController.Update)
	u.DELETE("/:userId", userController.Delete)
	u.GET("/", userController.FindAll)
	u.GET("/:userId", userController.FindById)

}
