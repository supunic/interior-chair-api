package controller

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ChairController struct {
	InputFactory      func(o port.ChairOutputPort, c port.ChairRepository) port.ChairInputPort
	OutputFactory     func(w http.ResponseWriter) port.ChairOutputPort
	RepositoryFactory func(c *gorm.DB) port.ChairRepository
	Conn              *gorm.DB
}

func (cc *ChairController) FindByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	outputPort := cc.OutputFactory(ctx.Response())
	repository := cc.RepositoryFactory(cc.Conn)
	inputPort := cc.InputFactory(outputPort, repository)
	inputPort.FindByID(id)

	return nil
}

func (cc *ChairController) Create(ctx echo.Context) error {
	var cid data.ChairInputData

	if err := ctx.Bind(&cid); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	outputPort := cc.OutputFactory(ctx.Response())
	repository := cc.RepositoryFactory(cc.Conn)
	inputPort := cc.InputFactory(outputPort, repository)
	inputPort.Create(&cid)

	return nil
}
