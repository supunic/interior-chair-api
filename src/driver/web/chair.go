package web

import (
	"app/adapter"
	"app/adapter/gateway"
	"app/adapter/presenter"
	"app/usecase/interactor"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initChairRouter(e *echo.Echo, conn *gorm.DB) *echo.Echo {
	ca := adapter.NewChairAdapter(
		interactor.NewChairInputPort,
		presenter.NewChairOutputPort,
		gateway.NewChairRepository,
		conn,
	)

	e.POST("/chair", ca.Create)
	e.GET("/chair/:id", ca.FindByID)

	return e
}
