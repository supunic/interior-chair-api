package router

import (
	"app/adapter/controller"
	"github.com/labstack/echo/v4"
)

type ChairAuthorRouter interface {
	Register()
}

type chairAuthorRouter struct {
	e  *echo.Echo
	cc controller.ChairAuthorController
}

func NewChairAuthorRouter(e *echo.Echo, cc controller.ChairAuthorController) ChairAuthorRouter {
	return &chairAuthorRouter{e: e, cc: cc}
}

func (cr *chairAuthorRouter) Register() {
	cr.e.GET("/chair-authors/:id", cr.cc.FetchAll)
}
