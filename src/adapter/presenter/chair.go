package presenter

import (
	"app/entity/model/chair"
	"app/usecase/port"
	"fmt"
	"net/http"
)

type ChairPresenter struct {
	w http.ResponseWriter
}

func NewChairOutputPort(w http.ResponseWriter) port.ChairOutputPort {
	return &ChairPresenter{w: w}
}

func (c *ChairPresenter) Render(chair *chair.Chair) {
	c.w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprint(c.w, chair.Name); err != nil {
		return
	}
}

func (c *ChairPresenter) RenderError(err error) {
	c.w.WriteHeader(http.StatusInternalServerError)

	if _, err := fmt.Fprint(c.w, err); err != nil {
		return
	}
}
