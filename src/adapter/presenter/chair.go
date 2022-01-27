package presenter

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type chairPresenter struct {
	c echo.Context
}

func NewChairOutputPort(c echo.Context) port.ChairOutputPort {
	return &chairPresenter{c: c}
}

func (cp *chairPresenter) Create(cod *data.ChairOutputData) {
	if err := cp.c.JSON(http.StatusOK, cod); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}

func (cp *chairPresenter) FetchAll(cod []*data.ChairOutputData) {
	if err := cp.c.JSON(http.StatusOK, cod); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}

func (cp *chairPresenter) FindByID(cod *data.ChairOutputData) {
	if err := cp.c.JSON(http.StatusOK, cod); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}

func (cp *chairPresenter) Error(err error) {
	if err := cp.c.JSON(http.StatusInternalServerError, err); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}
