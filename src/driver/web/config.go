package web

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func Init(e *echo.Echo, conn *gorm.DB) *echo.Echo {
	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Hello, World!")
	})

	initChairRouter(e, conn)

	return e
}
