package web

import (
	"app/adapter"
	"app/adapter/gateway"
	"app/adapter/presenter"
	"app/driver/web/router"
	"app/usecase/interactor"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// initApp ドメインごとにアプリ初期化
func initApp(e *echo.Echo, db *gorm.DB) {
	initChair(e, db)
}

func initChair(e *echo.Echo, db *gorm.DB) {
	ca := adapter.NewChairAdapter(
		interactor.NewChairInputPort,
		presenter.NewChairOutputPort,
		gateway.NewChairRepository,
		db,
	)
	cr := router.NewChairRouter(e, ca)
	cr.Register()
}