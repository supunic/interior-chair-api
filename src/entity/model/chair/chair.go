package chair

import (
	"app/entity/model/chairAuthor"
)

type Chair struct {
	ID            ID
	ChairAuthorID chairAuthor.ID
	Name          Name
	Feature       Feature
	Year          Year
	Image         Image
	ChairAuthor   chairAuthor.ChairAuthor
}

func NewChair(
	id uint,
	name string,
	feature string,
	year int,
	image string,
	chairAuthor *chairAuthor.ChairAuthor,
) (*Chair, error) {
	chairID, err := NewChairID(id)
	if err != nil {
		return nil, err
	}

	chairName, err := NewChairName(name)
	if err != nil {
		return nil, err
	}

	chairFeature, err := NewChairFeature(feature)
	if err != nil {
		return nil, err
	}

	chairYear, err := NewChairYear(year)
	if err != nil {
		return nil, err
	}

	chairImage, err := NewChairImage(image)
	if err != nil {
		return nil, err
	}

	return &Chair{
		ID:            *chairID,
		ChairAuthorID: chairAuthor.ID,
		Name:          *chairName,
		Feature:       *chairFeature,
		Year:          *chairYear,
		Image:         *chairImage,
		ChairAuthor:   *chairAuthor,
	}, nil
}
