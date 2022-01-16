package controller

import (
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
	chairName := ctx.Param("chairName")
	chairFeature := ctx.Param("chairFeature")
	chairImage := ctx.Param("chairImage")

	chairYear, err := strconv.Atoi(ctx.Param("chairYear"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	chairAuthorName := ctx.Param("chairAuthorName")
	chairAuthorDescription := ctx.Param("chairAuthorDescription")
	chairAuthorImage := ctx.Param("chairAuthorImage")

	chairAuthorBirthYear, err := strconv.Atoi(ctx.Param("chairAuthorBirthYear"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	chairAuthorDiedYear, err := strconv.Atoi(ctx.Param("chairAuthorDiedYear"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	outputPort := cc.OutputFactory(ctx.Response())
	repository := cc.RepositoryFactory(cc.Conn)
	inputPort := cc.InputFactory(outputPort, repository)
	inputPort.Create(
		chairName,
		chairFeature,
		chairYear,
		chairImage,
		chairAuthorName,
		chairAuthorDescription,
		chairAuthorBirthYear,
		chairAuthorDiedYear,
		chairAuthorImage,
	)

	return nil
}
