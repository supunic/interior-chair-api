package chairAuthor

type DiedYear int

func NewChairAuthorDiedYear(val int) (*DiedYear, error) {
	diedYear := DiedYear(val)
	return &diedYear, nil
}
