package controller

import (
	"go-crud-auth/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
	server *app.Server
}

func NewAuthController(server *app.Server) AuthController {
	return &AuthControllerImpl{server: server}
}

func (controller *AuthControllerImpl) Login(c echo.Context) error {
	return c.String(http.StatusOK, "AuthController - Login")
}

func (controller *AuthControllerImpl) Logout(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
