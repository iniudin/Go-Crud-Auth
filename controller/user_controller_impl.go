package controller

import (
	"go-crud-auth/helper"
	"go-crud-auth/model/web/requests"
	"go-crud-auth/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(c echo.Context) error {
	user := new(requests.UserCreate)
	err := c.Bind(user)
	helper.PanicError(err)

	userResponse := controller.UserService.Create(c.Request().Context(), *user)

	return c.JSON(http.StatusCreated, userResponse)
}

func (controller *UserControllerImpl) Update(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) Delete(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) FindById(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) FindAll(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}
