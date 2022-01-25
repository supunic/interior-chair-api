package interactor

import (
	"app/entity/model/chair"
	"app/entity/model/chairAuthor"
	"app/usecase/data"
	"app/usecase/port"
)

type chairInteractor struct {
	o port.ChairOutputPort
	r port.ChairRepository
}

func NewChairInputPort(o port.ChairOutputPort, r port.ChairRepository) port.ChairInputPort {
	return &chairInteractor{o: o, r: r}
}

func (ci *chairInteractor) Create(cid *data.ChairInputData) {
	newChairAuthor, err := chairAuthor.NewChairAuthor(
		cid.Author.Name,
		cid.Author.Description,
		cid.Author.BirthYear,
		cid.Author.DiedYear,
		cid.Author.Image,
	)

	if err != nil {
		ci.o.OutputError(err)
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
		ci.o.OutputError(err)
		return
	}

	c, err := ci.r.Create(newChair)

	if err != nil {
		ci.o.OutputError(err)
		return
	}

	cod := data.NewChairOutputData(c)

	ci.o.Output(cod)
}

func (ci *chairInteractor) FindByID(id uint) {
	chairID, err := chair.NewChairID(id)

	if err != nil {
		ci.o.OutputError(err)
		return
	}

	c, err := ci.r.FindByID(chairID)

	if err != nil {
		ci.o.OutputError(err)
		return
	}

	cod := data.NewChairOutputData(c)

	ci.o.Output(cod)
}
