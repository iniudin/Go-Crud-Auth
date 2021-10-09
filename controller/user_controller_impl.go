package controller

import (
	"github.com/labstack/echo/v4"
)

type UserControllerImpl struct {
}

func NewUserController() UserController {
	return &UserControllerImpl{}
}

func (controller *UserControllerImpl) Create(c echo.Context) error {
	panic("not implemented") // TODO: Implement
}

func (controller *UserControllerImpl) Read(c echo.Context) error {
	panic("not implemented") // TODO: Implement
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
