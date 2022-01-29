package chairAuthor

type ChairAuthor struct {
	ID          ID
	Name        Name
	Description Description
	BirthYear   BirthYear
	DiedYear    DiedYear
	Image       Image
}

func NewChairAuthor(
	name string,
	description string,
	birthYear int,
	diedYear int,
	image string,
) (*ChairAuthor, error) {
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
		Name:        *chairAuthorName,
		Description: *chairAuthorDescription,
		BirthYear:   *chairAuthorBirthYear,
		DiedYear:    *chairAuthorDiedYear,
		Image:       *chairAuthorImage,
	}, nil
}

func NewChairAuthorWithID(
	id uint,
	name string,
	description string,
	birthYear int,
	diedYear int,
	image string,
) (*ChairAuthor, error) {
	newChairAuthor, err := NewChairAuthor(name, description, birthYear, diedYear, image)

	if err != nil {
		return nil, err
	}

	chairID, err := NewChairAuthorID(id)

	if err != nil {
		return nil, err
	}

	newChairAuthor.ID = *chairID

	return newChairAuthor, nil
}
