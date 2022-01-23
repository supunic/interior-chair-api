package data

type ChairInputData struct {
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

//func NewChairOutputData(c *chair.Chair) (*ChairOutputData, error) {
//	return &ChairOutputData{
//		ID: c.ID.value(),
//	}, nil
//}
