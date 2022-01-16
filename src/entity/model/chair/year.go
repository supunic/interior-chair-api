package chair

type Year int

func NewChairYear(val int) (*Year, error) {
	year := Year(val)
	return &year, nil
}
