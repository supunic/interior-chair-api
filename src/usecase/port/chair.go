package port

import (
	"app/entity/model/chair"
	"app/usecase/data"
)

type ChairInputPort interface {
	Create(chairInputData *data.ChairCreateInputData)
	FetchAll()
	FindByID(id uint)
	Update(chairInputData *data.ChairUpdateInputData)
}

type ChairOutputPort interface {
	Create(chairOutputData *data.ChairOutputData)
	FetchAll(chairOutputDataList []*data.ChairOutputData)
	FindByID(chairOutputData *data.ChairOutputData)
	Update(chairOutputData *data.ChairOutputData)
	Error(err error)
}

type ChairRepository interface {
	Create(chair *chair.Chair) (*chair.Chair, error)
	FetchAll() ([]*chair.Chair, error)
	FindByID(id *chair.ID) (*chair.Chair, error)
	Update(chair *chair.Chair) (*chair.Chair, error)
}
