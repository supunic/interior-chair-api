package data

import "app/entity/model/chair"

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
	Author  AuthorOutputData `json:"author,omitempty"`
}

func NewChairOutputData(c *chair.Chair) *ChairOutputData {
	return &ChairOutputData{
		ID:      c.ID.Value(),
		Name:    c.Name.Value(),
		Feature: c.Feature.Value(),
		Year:    c.Year.Value(),
		Image:   c.Image.Value(),
		Author: AuthorOutputData{
			ID:          c.Author.ID.Value(),
			Name:        c.Author.Name.Value(),
			Description: c.Author.Name.Value(),
			BirthYear:   c.Author.BirthYear.Value(),
			DiedYear:    c.Author.DiedYear.Value(),
			Image:       c.Author.Image.Value(),
		},
	}
}
