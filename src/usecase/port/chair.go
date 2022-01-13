package port

import (
	"app/entity/model"
)

type ChairInputPort interface {
	FindByID(id int)
}

type ChairOutputPort interface {
	Render(chair *model.Chair)
	RenderError(err error)
}

type ChairRepository interface {
	FindByID(id int) (*model.Chair, error)
}
