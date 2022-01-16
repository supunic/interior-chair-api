package chair

import "app/entity/model/chairAuthor"

type Chair struct {
	ID      ID                      `json:"id"`
	Name    Name                    `json:"name"`
	Feature Feature                 `json:"feature"`
	Year    Year                    `json:"year"`
	Image   Image                   `json:"image"`
	Author  chairAuthor.ChairAuthor `json:"chairAuthor"`
}

func NewChair(
	id int,
	name string,
	feature string,
	year int,
	image string,
	chairAuthor *chairAuthor.ChairAuthor,
) (*Chair, error) {
	chairId, err := NewChairID(id)
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
		ID:      *chairId,
		Name:    *chairName,
		Feature: *chairFeature,
		Year:    *chairYear,
		Image:   *chairImage,
		Author:  *chairAuthor,
	}, nil
}
