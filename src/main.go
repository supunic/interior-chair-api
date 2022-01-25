package main

import (
	"app/driver/db"
	"app/driver/web"
	"github.com/labstack/echo/v4"
)

func main() {
	web.Serve(echo.New(), db.New())
}
