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
	chairNameObj, err := chair.NewChairName(chairName)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairFeatureObj, err := chair.NewChairFeature(chairFeature)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairYearObj, err := chair.NewChairYear(chairYear)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairImageObj, err := chair.NewChairImage(chairImage)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairAuthorNameObj, err := chairAuthor.NewChairAuthorName(chairAuthorName)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairAuthorDescriptionObj, err := chairAuthor.NewChairAuthorDescription(chairAuthorDescription)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairAuthorBirthYearObj, err := chairAuthor.NewChairAuthorBirthYear(chairAuthorBirthYear)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairAuthorDiedYearObj, err := chairAuthor.NewChairAuthorDiedYear(chairAuthorDiedYear)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	chairAuthorImageObj, err := chairAuthor.NewChairAuthorImage(chairAuthorImage)
	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	c, err := ci.Repository.Create(
		chairNameObj,
		chairFeatureObj,
		chairYearObj,
		chairImageObj,
		chairAuthorNameObj,
		chairAuthorDescriptionObj,
		chairAuthorBirthYearObj,
		chairAuthorDiedYearObj,
		chairAuthorImageObj,
	)

	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	ci.OutputPort.Render(c)
}
