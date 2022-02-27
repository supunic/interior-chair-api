package data

type AuthorInputData struct {
	ID          uint   `form:"authorId"  param:"id"`
	Name        string `form:"authorName"`
	Description string `form:"authorDescription"`
	BirthYear   int    `form:"authorBirthYear"`
	DiedYear    int    `form:"authorDiedYear"`
	Image       string `form:"authorImage"`
}

type AuthorOutputData struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	BirthYear   int    `json:"birthYear"`
	DiedYear    int    `json:"diedYear"`
	Image       string `json:"image"`
}

type AuthorRepositoryData struct {
	ID          uint   `gorm:""`
	Name        string `gorm:""`
	Description string `gorm:""`
	BirthYear   int    `gorm:""`
	DiedYear    int    `gorm:""`
	Image       string `gorm:""`
}

func (*AuthorRepositoryData) TableName() string {
	return "authors"
}

func NewAuthorOutputData(ard *AuthorRepositoryData) *AuthorOutputData {
	return &AuthorOutputData{
		ID:          ard.ID,
		Name:        ard.Name,
		Description: ard.Description,
		BirthYear:   ard.BirthYear,
		DiedYear:    ard.DiedYear,
		Image:       ard.Image,
	}
}
