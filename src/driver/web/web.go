package web

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Serve web サーバ起動
func Serve(e *echo.Echo, db *gorm.DB) {
	initApp(e, db)
	initMiddleware(e)
	e.Logger.Fatal(e.Start(":1323"))
}
