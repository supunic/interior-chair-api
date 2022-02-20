package author

type Name string

func NewAuthorName(val string) (*Name, error) {
	name := Name(val)
	return &name, nil
}

func (v *Name) Value() string {
	return string(*v)
}
