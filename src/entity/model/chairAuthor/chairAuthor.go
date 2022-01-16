package chairAuthor

type ChairAuthor struct {
	ID          ID          `json:"id"`
	Name        Name        `json:"name"`
	Description Description `json:"description"`
	BirthYear   BirthYear   `json:"birthYear"`
	DiedYear    DiedYear    `json:"diedYear"`
	Image       Image       `json:"image"`
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
