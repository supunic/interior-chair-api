package router

import (
	"app/adapter/controller"
	"github.com/labstack/echo/v4"
)

type UserRouter interface {
	Register()
}

type userRouter struct {
	e  *echo.Echo
	cc controller.UserController
}

func NewUserRouter(e *echo.Echo, cc controller.UserController) UserRouter {
	return &userRouter{e: e, cc: cc}
}

func (cr *userRouter) Register() {}
