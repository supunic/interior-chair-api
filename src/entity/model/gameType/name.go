package gameType

type Name string

func NewGameTypeName(val string) (*Name, error) {
	name := Name(val)
	return &name, nil
}
