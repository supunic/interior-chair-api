package data

import (
	"app/entity/model/author"
	"app/entity/model/chair"
)

type ChairInputData struct {
	ID      uint   `form:"chairId" param:"id"`
	Name    string `form:"chairName"`
	Feature string `form:"chairFeature"`
	Year    int    `form:"chairYear"`
	Image   string `form:"chairImage"`
	Author  AuthorInputData
}

type ChairOutputData struct {
	ID      uint             `json:"id"`
	Name    string           `json:"name"`
	Feature string           `json:"feature"`
	Year    int              `json:"year"`
	Image   string           `json:"image"`
	Author  AuthorOutputData `json:"author"`
}

type ChairRepositoryData struct {
	ID       uint   `gorm:""`
	AuthorID uint   `gorm:""`
	Name     string `gorm:""`
	Feature  string `gorm:""`
	Year     int    `gorm:""`
	Image    string `gorm:""`
	Author   AuthorRepositoryData
}

func (*ChairRepositoryData) TableName() string {
	return "chairs"
}

func (crd *ChairRepositoryData) Fill(c *chair.Chair, a *author.Author) {
	crd.ID = c.ID.Value()
	crd.AuthorID = a.ID.Value()
	crd.Name = c.Name.Value()
	crd.Feature = c.Feature.Value()
	crd.Year = c.Year.Value()
	crd.Image = c.Image.Value()
	crd.Author.ID = a.ID.Value()
	crd.Author.Name = a.Name.Value()
	crd.Author.Description = a.Description.Value()
	crd.Author.BirthYear = a.BirthYear.Value()
	crd.Author.DiedYear = a.DiedYear.Value()
	crd.Author.Image = a.Image.Value()
}

func (crd *ChairRepositoryData) FillID(id *chair.ID) {
	crd.ID = id.Value()
}

func NewChairOutputData(crd *ChairRepositoryData) *ChairOutputData {
	return &ChairOutputData{
		ID:      crd.ID,
		Name:    crd.Name,
		Feature: crd.Feature,
		Year:    crd.Year,
		Image:   crd.Image,
		Author: AuthorOutputData{
			ID:          crd.Author.ID,
			Name:        crd.Author.Name,
			Description: crd.Author.Description,
			BirthYear:   crd.Author.BirthYear,
			DiedYear:    crd.Author.DiedYear,
			Image:       crd.Author.Image,
		},
	}
}
