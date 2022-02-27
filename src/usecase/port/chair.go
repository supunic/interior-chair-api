package port

import (
	"app/usecase/data"
)

type ChairInputPort interface {
	Create(chairInputData *data.ChairInputData)
	Delete(id uint)
	FetchAll()
	FindByID(id uint)
	Update(chairInputData *data.ChairInputData)
}

type ChairOutputPort interface {
	Create(chairOutputData *data.ChairOutputData)
	Delete()
	FetchAll(chairOutputDataList []*data.ChairOutputData)
	FindByID(chairOutputData *data.ChairOutputData)
	Update(chairOutputData *data.ChairOutputData)
	Error(err error)
}

type ChairRepository interface {
	Create(chairRepositoryData *data.ChairRepositoryData) (*data.ChairRepositoryData, error)
	Delete(chairRepositoryData *data.ChairRepositoryData) error
	FetchAll() ([]*data.ChairRepositoryData, error)
	FindByID(chairRepositoryData *data.ChairRepositoryData) (*data.ChairRepositoryData, error)
	Update(chairRepositoryData *data.ChairRepositoryData) (*data.ChairRepositoryData, error)
}
