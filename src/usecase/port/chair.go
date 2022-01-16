package port

import (
	"app/entity/model/chair"
)

type ChairInputPort interface {
	FindByID(id int)
	Create(
		chairName string,
		chairFeature string,
		chairYear int,
		chairImage string,
		chairAuthorName string,
		chairAuthorDescription string,
		chairAuthorBirthYear int,
		chairAuthorDiedYear int,
		chairAuthorImage string,
	)
}

type ChairOutputPort interface {
	Render(chair *chair.Chair)
	RenderError(err error)
}

type ChairRepository interface {
	FindByID(id *chair.ID) (*chair.Chair, error)
	Create(chair *chair.Chair) (*chair.Chair, error)
}
