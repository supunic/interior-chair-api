package interactor

import (
	"app/usecase/port"
)

type chairAuthorInteractor struct {
	o port.ChairAuthorOutputPort
	r port.ChairAuthorRepository
}

func NewChairAuthorInputPort(o port.ChairAuthorOutputPort, r port.ChairAuthorRepository) port.ChairAuthorInputPort {
	return &chairAuthorInteractor{o: o, r: r}
}
