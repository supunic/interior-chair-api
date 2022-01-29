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
		cid.Author.ID,
		cid.Author.Name,
		cid.Author.Description,
		cid.Author.BirthYear,
		cid.Author.DiedYear,
		cid.Author.Image,
	)

	if err != nil {
		ci.o.Error(err)
		return
	}

	newChair, err := chair.NewChair(
		cid.ID,
		cid.Name,
		cid.Feature,
		cid.Year,
		cid.Image,
		newChairAuthor,
	)

	if err != nil {
		ci.o.Error(err)
		return
	}

	c, err := ci.r.Create(newChair)

	if err != nil {
		ci.o.Error(err)
		return
	}

	cod := data.NewChairOutputData(c)

	ci.o.Create(cod)
}

func (ci *chairInteractor) Delete(id uint) {
	chairID, err := chair.NewChairID(id)

	if err != nil {
		ci.o.Error(err)
		return
	}

	if err := ci.r.Delete(chairID); err != nil {
		ci.o.Error(err)
		return
	}

	ci.o.Delete()
}

func (ci *chairInteractor) FetchAll() {
	var cod []*data.ChairOutputData

	c, err := ci.r.FetchAll()

	if err != nil {
		ci.o.Error(err)
		return
	}

	for _, cv := range c {
		cod = append(cod, data.NewChairOutputData(cv))
	}

	ci.o.FetchAll(cod)
}

func (ci *chairInteractor) FindByID(id uint) {
	chairID, err := chair.NewChairID(id)

	if err != nil {
		ci.o.Error(err)
		return
	}

	c, err := ci.r.FindByID(chairID)

	if err != nil {
		ci.o.Error(err)
		return
	}

	cod := data.NewChairOutputData(c)

	ci.o.FindByID(cod)
}

func (ci *chairInteractor) Update(cid *data.ChairInputData) {
	newChairAuthor, err := chairAuthor.NewChairAuthor(
		cid.Author.ID,
		cid.Author.Name,
		cid.Author.Description,
		cid.Author.BirthYear,
		cid.Author.DiedYear,
		cid.Author.Image,
	)

	if err != nil {
		ci.o.Error(err)
		return
	}

	newChair, err := chair.NewChair(
		cid.ID,
		cid.Name,
		cid.Feature,
		cid.Year,
		cid.Image,
		newChairAuthor,
	)

	if err != nil {
		ci.o.Error(err)
		return
	}

	uc, err := ci.r.Update(newChair)

	if err != nil {
		ci.o.Error(err)
		return
	}

	cod := data.NewChairOutputData(uc)

	ci.o.Update(cod)
}
