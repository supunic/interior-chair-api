package web

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"net/http"
	"os"
)

func Init(e *echo.Echo, conn *gorm.DB) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if _, err := fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody)); err != nil {
			return
		}
	}))

	e.GET("/", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Hello, World!")
	})

	initChairRouter(e, conn)

	return e
}
