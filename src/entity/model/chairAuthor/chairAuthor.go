package chairAuthor

type ChairAuthor struct {
	ID          ID
	Name        Name
	Description Description
	BirthYear   BirthYear
	DiedYear    DiedYear
	Image       Image
	Chairs      []Chair
}

func NewChairAuthor(
	id uint,
	name string,
	description string,
	birthYear int,
	diedYear int,
	image string,
) (*ChairAuthor, error) {
	chairAuthorID, err := NewChairAuthorID(id)
	if err != nil {
		return nil, err
	}

	chairAuthorName, err := NewChairAuthorName(name)
	if err != nil {
		return nil, err
	}

	chairAuthorDescription, err := NewChairAuthorDescription(description)
	if err != nil {
		return nil, err
	}

	chairAuthorBirthYear, err := NewChairAuthorBirthYear(birthYear)
	if err != nil {
		return nil, err
	}

	chairAuthorDiedYear, err := NewChairAuthorDiedYear(diedYear)
	if err != nil {
		return nil, err
	}

	chairAuthorImage, err := NewChairAuthorImage(image)
	if err != nil {
		return nil, err
	}

	return &ChairAuthor{
		ID:          *chairAuthorID,
		Name:        *chairAuthorName,
		Description: *chairAuthorDescription,
		BirthYear:   *chairAuthorBirthYear,
		DiedYear:    *chairAuthorDiedYear,
		Image:       *chairAuthorImage,
	}, nil
}
