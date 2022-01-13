package controller

import (
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type Chair struct {
	OutputFactory     func(w http.ResponseWriter) port.ChairOutputPort
	InputFactory      func(o port.ChairOutputPort, c port.ChairRepository) port.ChairInputPort
	RepositoryFactory func(c *gorm.DB) port.ChairRepository
	Conn              *gorm.DB
}

func (c *Chair) FindByID(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	outputPort := c.OutputFactory(ctx.Response())
	repository := c.RepositoryFactory(c.Conn)
	inputPort := c.InputFactory(outputPort, repository)
	inputPort.FindByID(id)

	return nil
}
