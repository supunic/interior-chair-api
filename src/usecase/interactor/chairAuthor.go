package interactor

import (
	"app/usecase/data"
	"app/usecase/port"
)

type chairAuthorInteractor struct {
	o port.ChairAuthorOutputPort
	r port.ChairAuthorRepository
}

func NewChairAuthorInputPort(o port.ChairAuthorOutputPort, r port.ChairAuthorRepository) port.ChairAuthorInputPort {
	return &chairAuthorInteractor{o: o, r: r}
}

func (ci *chairAuthorInteractor) FetchAll() {
	var cod []*data.ChairAuthorOutputData

	c, err := ci.r.FetchAll()

	if err != nil {
		ci.o.Error(err)
		return
	}

	for _, cv := range c {
		cod = append(cod, data.NewChairAuthorOutputData(cv))
	}

	ci.o.FetchAll(cod)
}
