package author

type DiedYear int

func NewAuthorDiedYear(val int) (*DiedYear, error) {
	diedYear := DiedYear(val)
	return &diedYear, nil
}

func (v *DiedYear) Value() int {
	return int(*v)
}
