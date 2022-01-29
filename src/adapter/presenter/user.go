package presenter

import (
	"app/usecase/port"
	"github.com/labstack/echo/v4"
	"net/http"
)

type userPresenter struct {
	c echo.Context
}

func NewUserOutputPort(c echo.Context) port.UserOutputPort {
	return &userPresenter{c: c}
}

func (cp *userPresenter) Error(err error) {
	if err := cp.c.JSON(http.StatusInternalServerError, err.Error()); err != nil {
		http.Error(cp.c.Response(), err.Error(), http.StatusInternalServerError)
	}
}
