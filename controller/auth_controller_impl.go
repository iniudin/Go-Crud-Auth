package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthControllerImpl struct {
}

func NewAuthController() AuthController {
	return &AuthControllerImpl{}
}

func (controller *AuthControllerImpl) Login(c echo.Context) error {
	return c.String(http.StatusOK, "AuthController - Login")
}

func (controller *AuthControllerImpl) Logout(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
