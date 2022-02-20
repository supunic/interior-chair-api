package author

type BirthYear int

func NewAuthorBirthYear(val int) (*BirthYear, error) {
	birthYear := BirthYear(val)
	return &birthYear, nil
}

func (v *BirthYear) Value() int {
	return int(*v)
}
