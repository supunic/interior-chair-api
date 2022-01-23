package chair

import "app/entity/model/chairAuthor"

type Chair struct {
	ID       ID                      `json:"id"`
	AuthorID chairAuthor.ID          `json:"authorId"`
	Name     Name                    `json:"name"`
	Feature  Feature                 `json:"feature"`
	Year     Year                    `json:"year"`
	Image    Image                   `json:"image"`
	Author   chairAuthor.ChairAuthor `json:"chairAuthor"`
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
		Name:    *chairName,
		Feature: *chairFeature,
		Year:    *chairYear,
		Image:   *chairImage,
		Author:  *chairAuthor,
	}, nil
}
