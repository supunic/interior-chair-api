package presenter

import (
	"app/usecase/data"
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type authorPresenter struct {
	c echo.Context
}

func NewAuthorOutputPort(c echo.Context) port.AuthorOutputPort {
	return &authorPresenter{c: c}
}

func (cp *authorPresenter) Error(err error) {
	if err := cp.c.JSON(http.StatusInternalServerError, err.Error()); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}

func (cp *authorPresenter) FetchAll(cod []*data.AuthorOutputData) {
	if err := cp.c.JSON(http.StatusOK, cod); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}
