package data

import (
	"app/entity/model/author"
)

type AuthorInputData struct {
	ID          uint   `form:"authorId"  param:"id"`
	Name        string `form:"authorName"`
	Description string `form:"authorDescription"`
	BirthYear   int    `form:"authorBirthYear"`
	DiedYear    int    `form:"authorDiedYear"`
	Image       string `form:"authorImage"`
}

type AuthorOutputData struct {
	ID          uint               `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	BirthYear   int                `json:"birthYear"`
	DiedYear    int                `json:"diedYear"`
	Image       string             `json:"image"`
	Chairs      []*ChairOutputData `json:"chairs,omitempty"`
}

func NewAuthorOutputData(ca *author.Author) *AuthorOutputData {
	var chairs []*ChairOutputData

	for _, c := range ca.Chairs {
		cod := &ChairOutputData{
			ID:      c.ID,
			Name:    c.Name,
			Feature: c.Feature,
			Year:    c.Year,
			Image:   c.Image,
			Author: AuthorOutputData{
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

	return &AuthorOutputData{
		ID:          ca.ID.Value(),
		Name:        ca.Name.Value(),
		Description: ca.Description.Value(),
		BirthYear:   ca.BirthYear.Value(),
		DiedYear:    ca.DiedYear.Value(),
		Image:       ca.Image.Value(),
		Chairs:      chairs,
	}
}
