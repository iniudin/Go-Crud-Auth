package controller

import (
	"go-crud-auth/helper"
	"go-crud-auth/model/web/requests"
	"go-crud-auth/service"
	"net/http"
	"strconv"

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
	user := new(requests.UserUpdate)
	err := c.Bind(user)
	helper.PanicError(err)

	userId, _ := strconv.Atoi(c.Param("userId"))

	userResponse := controller.UserService.Update(c.Request().Context(), *user, uint(userId))

	return c.JSON(http.StatusCreated, userResponse)
}

func (controller *UserControllerImpl) Delete(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("userId"))
	controller.UserService.Delete(c.Request().Context(), uint(userId))

	return c.JSON(http.StatusOK, "User deleted")
}

func (controller *UserControllerImpl) FindById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("userId"))
	userResponse := controller.UserService.FindById(c.Request().Context(), uint(id))

	return c.JSON(http.StatusOK, userResponse)
}

func (controller *UserControllerImpl) FindAll(c echo.Context) error {
	userResponse := controller.UserService.FindAll(c.Request().Context())

	return c.JSON(http.StatusOK, userResponse)
}
