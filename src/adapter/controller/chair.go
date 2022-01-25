package controller

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type ChairController interface {
	Create() error
	FindByID() error
}

type chairController struct {
	c echo.Context
	i port.ChairInputPort
}

func NewChairController(c echo.Context, i port.ChairInputPort) ChairController {
	return &chairController{c: c, i: i}
}

func (cc *chairController) Create() error {
	var cid data.ChairInputData

	if err := cc.c.Bind(&cid); err != nil {
		return cc.c.JSON(http.StatusBadRequest, err)
	}

	cc.i.Create(&cid)

	return nil
}

func (cc *chairController) FindByID() error {
	id, err := strconv.Atoi(cc.c.Param("id"))

	if err != nil {
		return cc.c.JSON(http.StatusBadRequest, err)
	}

	cc.i.FindByID(uint(id))

	return nil
}
