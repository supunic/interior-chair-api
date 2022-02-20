package router

import (
	"app/adapter/controller"
	"github.com/labstack/echo/v4"
)

type AuthorRouter interface {
	Register()
}

type authorRouter struct {
	e  *echo.Echo
	cc controller.AuthorController
}

func NewAuthorRouter(e *echo.Echo, cc controller.AuthorController) AuthorRouter {
	return &authorRouter{e: e, cc: cc}
}

func (cr *authorRouter) Register() {
	cr.e.GET("/authors", cr.cc.FetchAll)
}
