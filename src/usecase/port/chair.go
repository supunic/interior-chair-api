package port

import (
	"app/entity/model/chair"
)

type ChairInputPort interface {
	FindByID(id int)
}

type ChairOutputPort interface {
	Render(chair *chair.Chair)
	RenderError(err error)
}

type ChairRepository interface {
	FindByID(id *chair.ID) (*chair.Chair, error)
}
