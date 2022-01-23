package data

type ChairAuthorInputData struct {
	Name        string `form:"chairAuthorName"`
	Description string `form:"chairAuthorDescription"`
	BirthYear   int    `form:"chairAuthorBirthYear"`
	DiedYear    int    `form:"chairAuthorDiedYear"`
	Image       string `form:"chairAuthorImage"`
}
