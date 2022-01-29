package interactor

import (
	"app/usecase/port"
)

type userInteractor struct {
	o port.UserOutputPort
	r port.UserRepository
}

func NewUserInputPort(o port.UserOutputPort, r port.UserRepository) port.UserInputPort {
	return &userInteractor{o: o, r: r}
}
