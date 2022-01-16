package chairAuthor

type Name string

func NewChairAuthorName(val string) (*Name, error) {
	name := Name(val)
	return &name, nil
}
