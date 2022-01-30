package data

import (
	"app/entity/model/chairAuthor"
)

type ChairAuthorInputData struct {
	ID          uint   `form:"chairAuthorId"  param:"id"`
	Name        string `form:"chairAuthorName"`
	Description string `form:"chairAuthorDescription"`
	BirthYear   int    `form:"chairAuthorBirthYear"`
	DiedYear    int    `form:"chairAuthorDiedYear"`
	Image       string `form:"chairAuthorImage"`
}

type ChairAuthorOutputData struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	BirthYear   int                `json:"birthYear"`
	DiedYear    int                `json:"diedYear"`
	Image       string             `json:"image"`
	Chairs      []*ChairOutputData `json:"chairs,omitempty"`
}

func NewChairAuthorOutputData(ca *chairAuthor.ChairAuthor) *ChairAuthorOutputData {
	var chairs []*ChairOutputData

	for _, c := range ca.Chairs {
		cod := &ChairOutputData{
			ID:      c.ID,
			Name:    c.Name,
			Feature: c.Feature,
			Year:    c.Year,
			Image:   c.Image,
			Author: ChairAuthorOutputData{
				ID:          ca.ID.Value(),
				Name:        ca.Name.Value(),
				Description: ca.Description.Value(),
				BirthYear:   ca.BirthYear.Value(),
				DiedYear:    ca.DiedYear.Value(),
				Image:       ca.Image.Value(),
			},
		}
		chairs = append(chairs, cod)
	}

	return &ChairAuthorOutputData{
		ID:          ca.ID.Value(),
		Name:        ca.Name.Value(),
		Description: ca.Description.Value(),
		BirthYear:   ca.BirthYear.Value(),
		DiedYear:    ca.DiedYear.Value(),
		Image:       ca.Image.Value(),
		Chairs:      chairs,
	}
}
