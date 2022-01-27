package router

import (
	"app/adapter"
	"github.com/labstack/echo/v4"
)

type ChairRouter interface {
	Register()
}

type chairRouter struct {
	e  *echo.Echo
	ca adapter.ChairAdapter
}

func NewChairRouter(e *echo.Echo, ca adapter.ChairAdapter) ChairRouter {
	return &chairRouter{e: e, ca: ca}
}

func (cr *chairRouter) Register() {
	cr.e.POST("/chairs", cr.ca.Create)
	cr.e.GET("/chairs", cr.ca.FetchAll)
	cr.e.GET("/chairs/:id", cr.ca.FindByID)
}
