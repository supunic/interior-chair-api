package chairAuthor

type BirthYear int

func NewChairAuthorBirthYear(val int) (*BirthYear, error) {
	birthYear := BirthYear(val)
	return &birthYear, nil
}
