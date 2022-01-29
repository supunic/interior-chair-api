package router

import (
	"app/adapter/controller"
	"github.com/labstack/echo/v4"
)

type ChairRouter interface {
	Register()
}

type chairRouter struct {
	e  *echo.Echo
	cc controller.ChairController
}

func NewChairRouter(e *echo.Echo, cc controller.ChairController) ChairRouter {
	return &chairRouter{e: e, cc: cc}
}

func (cr *chairRouter) Register() {
	cr.e.POST("/chairs", cr.cc.Create)
	cr.e.DELETE("/chairs", cr.cc.Delete)
	cr.e.GET("/chairs", cr.cc.FetchAll)
	cr.e.GET("/chairs/:id", cr.cc.FindByID)
	cr.e.PUT("/chairs", cr.cc.Update)
}
