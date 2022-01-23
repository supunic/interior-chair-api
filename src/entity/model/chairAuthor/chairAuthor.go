package chairAuthor

import "app/entity/builder"

type ChairAuthor struct {
	id          ID
	name        Name
	description Description
	birthYear   BirthYear
	diedYear    DiedYear
	image       Image
}

func (c *ChairAuthor) GetID() *ID                   { return &c.id }
func (c *ChairAuthor) GetName() *Name               { return &c.name }
func (c *ChairAuthor) GetDescription() *Description { return &c.description }
func (c *ChairAuthor) GetBirthYear() *BirthYear     { return &c.birthYear }
func (c *ChairAuthor) GetDiedYear() *DiedYear       { return &c.diedYear }
func (c *ChairAuthor) GetImage() *Image             { return &c.image }

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
		name:        *chairAuthorName,
		description: *chairAuthorDescription,
		birthYear:   *chairAuthorBirthYear,
		diedYear:    *chairAuthorDiedYear,
		image:       *chairAuthorImage,
	}, nil
}

func (c *ChairAuthor) RepositoryData(cn builder.ChairAuthorNotification) *builder.ChairAuthor {
	cb := cn.Build()

	cb.Name = c.name.Value()
	cb.Description = c.description.Value()
	cb.BirthYear = c.birthYear.Value()
	cb.DiedYear = c.diedYear.Value()
	cb.Image = c.image.Value()

	return cb
}
