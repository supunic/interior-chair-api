package data

type ChairAuthorInputData struct {
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
