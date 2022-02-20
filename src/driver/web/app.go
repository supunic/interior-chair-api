package web

import (
	"app/adapter/controller"
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
	initAuthor(e, db)
	initUser(e, db)
}

func initChair(e *echo.Echo, db *gorm.DB) {
	cc := controller.NewChairController(
		interactor.NewChairInputPort,
		presenter.NewChairOutputPort,
		gateway.NewChairRepository,
		db,
	)
	cr := router.NewChairRouter(e, cc)
	cr.Register()
}

func initAuthor(e *echo.Echo, db *gorm.DB) {
	cc := controller.NewAuthorController(
		interactor.NewAuthorInputPort,
		presenter.NewAuthorOutputPort,
		gateway.NewAuthorRepository,
		db,
	)
	cr := router.NewAuthorRouter(e, cc)
	cr.Register()
}

func initUser(e *echo.Echo, db *gorm.DB) {
	cc := controller.NewUserController(
		interactor.NewUserInputPort,
		presenter.NewUserOutputPort,
		gateway.NewUserRepository,
		db,
	)
	cr := router.NewUserRouter(e, cc)
	cr.Register()
}
