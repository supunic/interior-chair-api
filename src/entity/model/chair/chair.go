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
	name string,
	feature string,
	year int,
	image string,
	chairAuthor *chairAuthor.ChairAuthor,
) (*Chair, error) {
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
		Name:        *chairName,
		Feature:     *chairFeature,
		Year:        *chairYear,
		Image:       *chairImage,
		ChairAuthor: *chairAuthor,
	}, nil
}

func NewChairWithID(
	id uint,
	chairAuthorId *chairAuthor.ID,
	name string,
	feature string,
	year int,
	image string,
	chairAuthor *chairAuthor.ChairAuthor,
) (*Chair, error) {
	newChair, err := NewChair(name, feature, year, image, chairAuthor)

	if err != nil {
		return nil, err
	}

	chairID, err := NewChairID(id)

	if err != nil {
		return nil, err
	}

	newChair.ID = *chairID
	newChair.ChairAuthorID = *chairAuthorId

	return newChair, nil
}
