package chair

type Chair struct {
	ID      ID
	Name    Name
	Feature Feature
	Year    Year
	Image   Image
}

func NewChair(
	id uint,
	name string,
	feature string,
	year int,
	image string,
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
		ID:      *chairID,
		Name:    *chairName,
		Feature: *chairFeature,
		Year:    *chairYear,
		Image:   *chairImage,
	}, nil
}
