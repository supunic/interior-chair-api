package main

import (
	"app/driver/db"
	"app/driver/web"
	"github.com/labstack/echo/v4"
)

func main() {
	conn := db.Init()
	e := web.Init(echo.New(), conn)
	e.Logger.Fatal(e.Start(":1323"))
}
