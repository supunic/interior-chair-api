package interactor

import (
	"app/entity/model/chair"
	"app/usecase/port"
)

type ChairInteractor struct {
	OutputPort port.ChairOutputPort
	Repository port.ChairRepository
}

func NewChairInputPort(outputPort port.ChairOutputPort, repository port.ChairRepository) port.ChairInputPort {
	return &ChairInteractor{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (ci *ChairInteractor) FindByID(id int) {
	chairID, err := chair.NewChairID(id)

	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	c, err := ci.Repository.FindByID(chairID)

	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	ci.OutputPort.Render(c)
}
