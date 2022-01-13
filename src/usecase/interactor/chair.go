package interactor

import (
	"app/usecase/port"
)

type Chair struct {
	OutputPort port.ChairOutputPort
	Repository port.ChairRepository
}

func NewChairInputPort(outputPort port.ChairOutputPort, repository port.ChairRepository) port.ChairInputPort {
	return &Chair{
		OutputPort: outputPort,
		Repository: repository,
	}
}

func (c *Chair) FindByID(id int) {
	chair, err := c.Repository.FindByID(id)

	if err != nil {
		c.OutputPort.RenderError(err)
		return
	}

	c.OutputPort.Render(chair)
}
