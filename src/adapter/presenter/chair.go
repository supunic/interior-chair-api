package presenter

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ChairPresenter struct {
	c echo.Context
}

func NewChairOutputPort(c echo.Context) port.ChairOutputPort {
	return &ChairPresenter{c: c}
}

func (cp *ChairPresenter) Output(cod *data.ChairOutputData) {
	if err := cp.c.JSON(http.StatusOK, cod); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}

func (cp *ChairPresenter) OutputError(err error) {
	if err := cp.c.JSON(http.StatusInternalServerError, err); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}
