package author

type Author struct {
	ID          ID
	Name        Name
	Description Description
	BirthYear   BirthYear
	DiedYear    DiedYear
	Image       Image
}

func NewAuthor(
	id uint,
	name string,
	description string,
	birthYear int,
	diedYear int,
	image string,
) (*Author, error) {
	authorID, err := NewAuthorID(id)
	if err != nil {
		return nil, err
	}

	authorName, err := NewAuthorName(name)
	if err != nil {
		return nil, err
	}

	authorDescription, err := NewAuthorDescription(description)
	if err != nil {
		return nil, err
	}

	authorBirthYear, err := NewAuthorBirthYear(birthYear)
	if err != nil {
		return nil, err
	}

	authorDiedYear, err := NewAuthorDiedYear(diedYear)
	if err != nil {
		return nil, err
	}

	authorImage, err := NewAuthorImage(image)
	if err != nil {
		return nil, err
	}

	return &Author{
		ID:          *authorID,
		Name:        *authorName,
		Description: *authorDescription,
		BirthYear:   *authorBirthYear,
		DiedYear:    *authorDiedYear,
		Image:       *authorImage,
	}, nil
}
