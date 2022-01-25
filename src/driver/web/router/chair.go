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
	cr.e.POST("/chair", cr.ca.Create)
	cr.e.GET("/chair/:id", cr.ca.FindByID)
}
