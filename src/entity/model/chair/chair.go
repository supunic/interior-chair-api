package chair

import (
	"app/entity/builder"
	"app/entity/model/chairAuthor"
)

type Chair struct {
	id      ID
	name    Name
	feature Feature
	year    Year
	image   Image
	author  chairAuthor.ChairAuthor
}

func (c *Chair) GetID() *ID                          { return &c.id }
func (c *Chair) GetName() *Name                      { return &c.name }
func (c *Chair) GetFeature() *Feature                { return &c.feature }
func (c *Chair) GetYear() *Year                      { return &c.year }
func (c *Chair) GetImage() *Image                    { return &c.image }
func (c *Chair) GetAuthor() *chairAuthor.ChairAuthor { return &c.author }

func (c *Chair) SetID(id int) error {
	chairID, err := NewChairID(id)
	if err != nil {
		return err
	}
	c.id = *chairID
	return nil
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
		name:    *chairName,
		feature: *chairFeature,
		year:    *chairYear,
		image:   *chairImage,
		author:  *chairAuthor,
	}, nil
}

func (c *Chair) RepositoryData(cn builder.ChairNotification) *builder.Chair {
	cb := cn.Build()

	cb.Name = c.GetName().Value()
	cb.Feature = c.GetFeature().Value()
	cb.Year = c.GetYear().Value()
	cb.Image = c.GetImage().Value()
	cb.Author.Name = c.GetAuthor().GetName().Value()
	cb.Author.Description = c.GetAuthor().GetDescription().Value()
	cb.Author.BirthYear = c.GetAuthor().GetBirthYear().Value()
	cb.Author.DiedYear = c.GetAuthor().GetDiedYear().Value()
	cb.Author.Image = c.GetAuthor().GetImage().Value()

	return cb
}
