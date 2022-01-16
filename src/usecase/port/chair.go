package port

import (
	"app/entity/model/chair"
	"app/entity/model/chairAuthor"
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
	Create(
		chairName *chair.Name,
		chairFeature *chair.Feature,
		chairYear *chair.Year,
		chairImage *chair.Image,
		chairAuthorName *chairAuthor.Name,
		chairAuthorDescription *chairAuthor.Description,
		chairAuthorBirthYear *chairAuthor.BirthYear,
		chairAuthorDiedYear *chairAuthor.DiedYear,
		chairAuthorImage *chairAuthor.Image,
	) (*chair.Chair, error)
}
