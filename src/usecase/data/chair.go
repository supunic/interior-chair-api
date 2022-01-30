package data

import "app/entity/model/chair"

type ChairInputData struct {
	ID      uint   `form:"chairId" param:"id"`
	Name    string `form:"chairName"`
	Feature string `form:"chairFeature"`
	Year    int    `form:"chairYear"`
	Image   string `form:"chairImage"`
	Author  ChairAuthorInputData
}

type ChairOutputData struct {
	ID      uint                  `json:"id"`
	Name    string                `json:"name"`
	Feature string                `json:"feature"`
	Year    int                   `json:"year"`
	Image   string                `json:"image"`
	Author  ChairAuthorOutputData `json:"author,omitempty"`
}

func NewChairOutputData(c *chair.Chair) *ChairOutputData {
	return &ChairOutputData{
		ID:      c.ID.Value(),
		Name:    c.Name.Value(),
		Feature: c.Feature.Value(),
		Year:    c.Year.Value(),
		Image:   c.Image.Value(),
		Author: ChairAuthorOutputData{
			ID:          c.ChairAuthor.ID.Value(),
			Name:        c.ChairAuthor.Name.Value(),
			Description: c.ChairAuthor.Name.Value(),
			BirthYear:   c.ChairAuthor.BirthYear.Value(),
			DiedYear:    c.ChairAuthor.DiedYear.Value(),
			Image:       c.ChairAuthor.Image.Value(),
		},
	}
}
