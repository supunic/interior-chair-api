package interactor

import (
	"app/usecase/data"
	"app/usecase/port"
)

type authorInteractor struct {
	o port.AuthorOutputPort
	r port.AuthorRepository
}

func NewAuthorInputPort(o port.AuthorOutputPort, r port.AuthorRepository) port.AuthorInputPort {
	return &authorInteractor{o: o, r: r}
}

func (ci *authorInteractor) FetchAll() {
	var cod []*data.AuthorOutputData

	c, err := ci.r.FetchAll()

	if err != nil {
		ci.o.Error(err)
		return
	}

	for _, cv := range c {
		cod = append(cod, data.NewAuthorOutputData(cv))
	}

	ci.o.FetchAll(cod)
}
