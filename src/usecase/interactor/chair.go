package interactor

import (
	"app/entity/model/chair"
	"app/entity/model/chairAuthor"
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

func (ci *ChairInteractor) Create(
	chairName string,
	chairFeature string,
	chairYear int,
	chairImage string,
	chairAuthorName string,
	chairAuthorDescription string,
	chairAuthorBirthYear int,
	chairAuthorDiedYear int,
	chairAuthorImage string,
) {
	newChairAuthor, err := chairAuthor.NewChairAuthor(
		chairAuthorName,
		chairAuthorDescription,
		chairAuthorBirthYear,
		chairAuthorDiedYear,
		chairAuthorImage,
	)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	newChair, err := chair.NewChair(
		chairName,
		chairFeature,
		chairYear,
		chairImage,
		newChairAuthor,
	)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	c, err := ci.Repository.Create(newChair)

	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	ci.OutputPort.Render(c)
}
