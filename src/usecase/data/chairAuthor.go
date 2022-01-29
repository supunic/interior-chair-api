package data

import "app/entity/model/chairAuthor"

type ChairAuthorCreateInputData struct {
	ID          uint   `form:"chairAuthorId"`
	Name        string `form:"chairAuthorName"`
	Description string `form:"chairAuthorDescription"`
	BirthYear   int    `form:"chairAuthorBirthYear"`
	DiedYear    int    `form:"chairAuthorDiedYear"`
	Image       string `form:"chairAuthorImage"`
}

type ChairAuthorUpdateInputData struct {
	ID          uint   `form:"chairAuthorId"`
	Name        string `form:"chairAuthorName"`
	Description string `form:"chairAuthorDescription"`
	BirthYear   int    `form:"chairAuthorBirthYear"`
	DiedYear    int    `form:"chairAuthorDiedYear"`
	Image       string `form:"chairAuthorImage"`
}

type ChairAuthorOutputData struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BirthYear   int    `json:"birthYear"`
	DiedYear    int    `json:"diedYear"`
	Image       string `json:"image"`
}

func NewChairAuthorOutputData(c *chairAuthor.ChairAuthor) *ChairAuthorOutputData {
	return &ChairAuthorOutputData{
		ID:          c.ID.Value(),
		Name:        c.Name.Value(),
		Description: c.Description.Value(),
		BirthYear:   c.BirthYear.Value(),
		DiedYear:    c.DiedYear.Value(),
		Image:       c.Image.Value(),
	}
}
