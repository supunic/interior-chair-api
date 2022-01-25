package port

import (
	"app/entity/model/chair"
	"app/usecase/data"
)

type ChairInputPort interface {
	FindByID(id uint)
	Create(chairInputData *data.ChairInputData)
}

type ChairOutputPort interface {
	Output(cod *data.ChairOutputData)
	OutputError(err error)
}

type ChairRepository interface {
	FindByID(id *chair.ID) (*chair.Chair, error)
	Create(chair *chair.Chair) (*chair.Chair, error)
}
