package data

import "app/entity/model/chair"

type ChairInputData struct {
	ID      uint   `form:"chairId"`
	Name    string `form:"chairName"`
	Feature string `form:"chairFeature"`
	Year    int    `form:"chairYear"`
	Image   string `form:"chairImage"`
	Author  ChairAuthorInputData
}

type ChairOutputData struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Feature string `json:"feature"`
	Year    int    `json:"year"`
	Image   string `json:"image"`
	Author  ChairAuthorOutputData
}

func NewChairOutputData(c *chair.Chair) *ChairOutputData {
	return &ChairOutputData{
		ID:      c.ID.Value(),
		Name:    c.Name.Value(),
		Feature: c.Feature.Value(),
		Year:    c.Year.Value(),
		Image:   c.Image.Value(),
		Author:  *NewChairAuthorOutputData(&c.ChairAuthor),
	}
}
