package controller

import (
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ChairAuthorController interface{}

type newChairAuthorInputPort func(o port.ChairAuthorOutputPort, r port.ChairAuthorRepository) port.ChairAuthorInputPort
type newChairAuthorOutputPort func(c echo.Context) port.ChairAuthorOutputPort
type newChairAuthorRepository func(c *gorm.DB) port.ChairAuthorRepository

type chairAuthorController struct {
	newInputPort  newChairAuthorInputPort
	newOutputPort newChairAuthorOutputPort
	newRepository newChairAuthorRepository
	db            *gorm.DB
}

func NewChairAuthorController(
	i newChairAuthorInputPort,
	o newChairAuthorOutputPort,
	r newChairAuthorRepository,
	db *gorm.DB,
) ChairAuthorController {
	return &chairAuthorController{newInputPort: i, newOutputPort: o, newRepository: r, db: db}
}

func (cc *chairAuthorController) handler(c echo.Context) port.ChairAuthorInputPort {
	cop := cc.newOutputPort(c)
	cr := cc.newRepository(cc.db)
	cip := cc.newInputPort(cop, cr)

	return cip
}
