package presenter

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type chairAuthorPresenter struct {
	c echo.Context
}

func NewChairAuthorOutputPort(c echo.Context) port.ChairAuthorOutputPort {
	return &chairAuthorPresenter{c: c}
}

func (cp *chairAuthorPresenter) Error(err error) {
	if err := cp.c.JSON(http.StatusInternalServerError, err.Error()); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}

func (cp *chairAuthorPresenter) FetchAll(cod []*data.ChairAuthorOutputData) {
	if err := cp.c.JSON(http.StatusOK, cod); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}
