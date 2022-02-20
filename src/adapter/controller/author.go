package controller

import (
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthorController interface {
	FetchAll(c echo.Context) error
}

type newAuthorInputPort func(o port.AuthorOutputPort, r port.AuthorRepository) port.AuthorInputPort
type newAuthorOutputPort func(c echo.Context) port.AuthorOutputPort
type newAuthorRepository func(c *gorm.DB) port.AuthorRepository

type authorController struct {
	newInputPort  newAuthorInputPort
	newOutputPort newAuthorOutputPort
	newRepository newAuthorRepository
	db            *gorm.DB
}

func NewAuthorController(
	i newAuthorInputPort,
	o newAuthorOutputPort,
	r newAuthorRepository,
	db *gorm.DB,
) AuthorController {
	return &authorController{newInputPort: i, newOutputPort: o, newRepository: r, db: db}
}

func (cc *authorController) handler(c echo.Context) port.AuthorInputPort {
	cop := cc.newOutputPort(c)
	cr := cc.newRepository(cc.db)
	cip := cc.newInputPort(cop, cr)

	return cip
}

func (cc *authorController) FetchAll(c echo.Context) error {
	cc.handler(c).FetchAll()

	return nil
}
