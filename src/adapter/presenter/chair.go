package presenter

import (
	"app/entity/model"
	"app/usecase/port"
	"fmt"
	"net/http"
)

type Chair struct {
	w http.ResponseWriter
}

func NewChairOutputPort(w http.ResponseWriter) port.ChairOutputPort {
	return &Chair{w: w}
}

func (c *Chair) Render(chair *model.Chair) {
	c.w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprint(c.w, chair.Name); err != nil {
		return
	}
}

func (c *Chair) RenderError(err error) {
	c.w.WriteHeader(http.StatusInternalServerError)

	if _, err := fmt.Fprint(c.w, err); err != nil {
		return
	}
}
