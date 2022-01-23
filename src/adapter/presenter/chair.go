package presenter

import (
	"app/usecase/data"
	"app/usecase/port"
	"encoding/json"
	"fmt"
	"net/http"
)

type ChairPresenter struct {
	w http.ResponseWriter
}

func NewChairOutputPort(w http.ResponseWriter) port.ChairOutputPort {
	return &ChairPresenter{w: w}
}

func (c *ChairPresenter) Render(cod *data.ChairOutputData) {
	res, err := json.Marshal(cod)

	if err != nil {
		http.Error(c.w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.w.Header().Set("Content-Type", "application/json")

	if _, err := c.w.Write(res); err != nil {
		return
	}
}

func (c *ChairPresenter) RenderError(err error) {
	c.w.WriteHeader(http.StatusInternalServerError)

	if _, err := fmt.Fprint(c.w, err); err != nil {
		return
	}
}
