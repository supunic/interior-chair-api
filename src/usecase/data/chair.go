package data

type ChairInputData struct {
	Name    string `form:"chairName"`
	Feature string `form:"chairFeature"`
	Year    int    `form:"chairYear"`
	Image   string `form:"chairImage"`
	Author  ChairAuthorInputData
}
