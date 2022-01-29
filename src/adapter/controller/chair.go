package controller

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ChairController interface {
	Create(c echo.Context) error
	Delete(c echo.Context) error
	FetchAll(c echo.Context) error
	FindByID(c echo.Context) error
	Update(c echo.Context) error
}

type newChairInputPort func(o port.ChairOutputPort, r port.ChairRepository) port.ChairInputPort
type newChairOutputPort func(c echo.Context) port.ChairOutputPort
type newChairRepository func(c *gorm.DB) port.ChairRepository

type chairController struct {
	newInputPort  newChairInputPort
	newOutputPort newChairOutputPort
	newRepository newChairRepository
	db            *gorm.DB
}

func NewChairController(i newChairInputPort, o newChairOutputPort, r newChairRepository, db *gorm.DB) ChairController {
	return &chairController{newInputPort: i, newOutputPort: o, newRepository: r, db: db}
}

func (cc *chairController) Create(c echo.Context) error {
	var cid data.ChairCreateInputData

	if err := c.Bind(&cid); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cc.handler(c).Create(&cid)

	return nil
}

func (cc *chairController) Delete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cc.handler(c).Delete(uint(id))

	return nil
}

func (cc *chairController) FetchAll(c echo.Context) error {
	cc.handler(c).FetchAll()

	return nil
}

func (cc *chairController) FindByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cc.handler(c).FindByID(uint(id))

	return nil
}

func (cc *chairController) Update(c echo.Context) error {
	var cid data.ChairUpdateInputData

	if err := c.Bind(&cid); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	cc.handler(c).Update(&cid)

	return nil
}

func (cc *chairController) handler(c echo.Context) port.ChairInputPort {
	// context を受け取る度に生成しないといけない
	cop := cc.newOutputPort(c)
	cr := cc.newRepository(cc.db)
	cip := cc.newInputPort(cop, cr)

	return cip
}
