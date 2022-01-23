package interactor

import (
	"app/entity/model/chair"
	"app/entity/model/chairAuthor"
	"app/usecase/data"
	"app/usecase/port"
	"fmt"
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

	fmt.Println(c) // 修正必要
	//ci.OutputPort.Render(c)
}

func (ci *ChairInteractor) Create(cid *data.ChairInputData) {
	newChairAuthor, err := chairAuthor.NewChairAuthor(
		cid.Author.Name,
		cid.Author.Description,
		cid.Author.BirthYear,
		cid.Author.DiedYear,
		cid.Author.Image,
	)

	if err != nil {
		ci.OutputPort.RenderError(err)
		return
	}

	newChair, err := chair.NewChair(
		cid.Name,
		cid.Feature,
		cid.Year,
		cid.Image,
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

	cod := data.ChairOutputData{
		ID:      c.GetID().Value(),
		Name:    c.GetName().Value(),
		Feature: c.GetFeature().Value(),
		Year:    c.GetYear().Value(),
		Image:   c.GetImage().Value(),
		Author: data.ChairAuthorOutputData{
			ID:          c.GetAuthor().GetID().Value(),
			Name:        c.GetAuthor().GetName().Value(),
			Description: c.GetAuthor().GetDescription().Value(),
			BirthYear:   c.GetAuthor().GetBirthYear().Value(),
			DiedYear:    c.GetAuthor().GetDiedYear().Value(),
			Image:       c.GetAuthor().GetImage().Value(),
		},
	}

	ci.OutputPort.Render(&cod)
}
