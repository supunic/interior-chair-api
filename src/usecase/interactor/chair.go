package interactor

import (
	"app/entity/model/author"
	"app/entity/model/chair"
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
	var crd data.ChairRepositoryData

	na, err := author.NewAuthor(
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

	nc, err := chair.NewChair(
		cid.ID,
		cid.Name,
		cid.Feature,
		cid.Year,
		cid.Image,
	)

	if err != nil {
		ci.o.Error(err)
		return
	}

	crd.Fill(nc, na)

	c, err := ci.r.Create(&crd)

	if err != nil {
		ci.o.Error(err)
		return
	}

	cod := data.NewChairOutputData(c)

	ci.o.Create(cod)
}

func (ci *chairInteractor) Delete(id uint) {
	var crd data.ChairRepositoryData

	cid, err := chair.NewChairID(id)

	if err != nil {
		ci.o.Error(err)
		return
	}

	crd.FillID(cid)

	if err := ci.r.Delete(&crd); err != nil {
		ci.o.Error(err)
		return
	}

	ci.o.Delete()
}

func (ci *chairInteractor) FetchAll() {
	var cod []*data.ChairOutputData

	cs, err := ci.r.FetchAll()

	if err != nil {
		ci.o.Error(err)
		return
	}

	for _, c := range cs {
		cod = append(cod, data.NewChairOutputData(c))
	}

	ci.o.FetchAll(cod)
}

func (ci *chairInteractor) FindByID(id uint) {
	var crd data.ChairRepositoryData

	cid, err := chair.NewChairID(id)

	if err != nil {
		ci.o.Error(err)
		return
	}

	crd.FillID(cid)

	c, err := ci.r.FindByID(&crd)

	if err != nil {
		ci.o.Error(err)
		return
	}

	cod := data.NewChairOutputData(c)

	ci.o.FindByID(cod)
}

func (ci *chairInteractor) Update(cid *data.ChairInputData) {
	var crd data.ChairRepositoryData

	na, err := author.NewAuthor(
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

	nc, err := chair.NewChair(
		cid.ID,
		cid.Name,
		cid.Feature,
		cid.Year,
		cid.Image,
	)

	if err != nil {
		ci.o.Error(err)
		return
	}

	crd.Fill(nc, na)

	uc, err := ci.r.Update(&crd)

	if err != nil {
		ci.o.Error(err)
		return
	}

	cod := data.NewChairOutputData(uc)

	ci.o.Update(cod)
}
