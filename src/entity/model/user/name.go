package user

type Name string

func NewUserName(val string) (*Name, error) {
	name := Name(val)
	return &name, nil
}
