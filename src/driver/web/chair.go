package web

import (
	"app/adapter/controller"
	"app/adapter/gateway"
	"app/adapter/presenter"
	"app/usecase/interactor"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func initChairRouter(e *echo.Echo, conn *gorm.DB) *echo.Echo {
	chair := controller.Chair{
		OutputFactory:     presenter.NewChairOutputPort,
		InputFactory:      interactor.NewChairInputPort,
		RepositoryFactory: gateway.NewChairRepository,
		Conn:              conn,
	}

	e.GET("/chair/:id", chair.FindByID)

	return e
}
