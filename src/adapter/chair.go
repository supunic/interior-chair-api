package adapter

import (
	"app/adapter/controller"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ChairAdapter interface {
	Create(c echo.Context) error
	FindByID(c echo.Context) error
}

type chairAdapter struct {
	newInputPort  func(o port.ChairOutputPort, r port.ChairRepository) port.ChairInputPort
	newOutputPort func(c echo.Context) port.ChairOutputPort
	newRepository func(c *gorm.DB) port.ChairRepository
	db            *gorm.DB
}

func NewChairAdapter(
	i func(o port.ChairOutputPort, r port.ChairRepository) port.ChairInputPort,
	o func(c echo.Context) port.ChairOutputPort,
	r func(c *gorm.DB) port.ChairRepository,
	db *gorm.DB,
) ChairAdapter {
	return &chairAdapter{newInputPort: i, newOutputPort: o, newRepository: r, db: db}
}

func (ca *chairAdapter) Create(c echo.Context) error {
	if err := ca.handler(c).Create(); err != nil {
		return err
	}

	return nil
}

func (ca *chairAdapter) FindByID(c echo.Context) error {
	if err := ca.handler(c).FindByID(); err != nil {
		return err
	}

	return nil
}

func (ca *chairAdapter) handler(c echo.Context) controller.ChairController {
	// context を受け取る度に生成しないといけない
	cop := ca.newOutputPort(c)
	cr := ca.newRepository(ca.db)
	cip := ca.newInputPort(cop, cr)

	return controller.NewChairController(c, cip)
}
