package port

import (
	"app/entity/model/chair"
	"app/usecase/data"
)

type ChairInputPort interface {
	Create(chairInputData *data.ChairInputData)
	FetchAll()
	FindByID(id uint)
}

type ChairOutputPort interface {
	Create(cod *data.ChairOutputData)
	FetchAll(cod []*data.ChairOutputData)
	FindByID(cod *data.ChairOutputData)
	Error(err error)
}

type ChairRepository interface {
	Create(chair *chair.Chair) (*chair.Chair, error)
	FetchAll() ([]*chair.Chair, error)
	FindByID(id *chair.ID) (*chair.Chair, error)
}
