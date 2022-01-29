package controller

import (
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface{}

type newUserInputPort func(o port.UserOutputPort, r port.UserRepository) port.UserInputPort
type newUserOutputPort func(c echo.Context) port.UserOutputPort
type newUserRepository func(c *gorm.DB) port.UserRepository

type userController struct {
	newInputPort  newUserInputPort
	newOutputPort newUserOutputPort
	newRepository newUserRepository
	db            *gorm.DB
}

func NewUserController(
	i newUserInputPort,
	o newUserOutputPort,
	r newUserRepository,
	db *gorm.DB,
) UserController {
	return &userController{newInputPort: i, newOutputPort: o, newRepository: r, db: db}
}

func (cc *userController) handler(c echo.Context) port.UserInputPort {
	cop := cc.newOutputPort(c)
	cr := cc.newRepository(cc.db)
	cip := cc.newInputPort(cop, cr)

	return cip
}
