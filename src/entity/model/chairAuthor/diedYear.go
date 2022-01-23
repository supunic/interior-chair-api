package chairAuthor

type DiedYear int

func NewChairAuthorDiedYear(val int) (*DiedYear, error) {
	diedYear := DiedYear(val)
	return &diedYear, nil
}

func (v *DiedYear) Value() int {
	return int(*v)
}
