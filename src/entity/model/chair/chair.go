package chair

import "app/entity/model/author"

type Chair struct {
	ID       ID
	AuthorID author.ID
	Name     Name
	Feature  Feature
	Year     Year
	Image    Image
}

func NewChair(
	id uint,
	authorId uint,
	name string,
	feature string,
	year int,
	image string,
) (*Chair, error) {
	chairID, err := NewChairID(id)
	if err != nil {
		return nil, err
	}

	authorID, err := author.NewAuthorID(authorId)
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
		ID:       *chairID,
		AuthorID: *authorID,
		Name:     *chairName,
		Feature:  *chairFeature,
		Year:     *chairYear,
		Image:    *chairImage,
	}, nil
}
