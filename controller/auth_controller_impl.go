package controller

import (
	"go-crud-auth/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	Server *app.Server
}

func NewAuthController(server *app.Server) AuthController {
	return &AuthControllerImpl{Server: server}
}

func (controller *AuthControllerImpl) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "AuthController - Login")
}

func (controller *AuthControllerImpl) Logout(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
