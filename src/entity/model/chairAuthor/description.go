package chairAuthor

type Description string

func NewChairAuthorDescription(val string) (*Description, error) {
	description := Description(val)
	return &description, nil
}

func (v *Description) Value() string {
	return string(*v)
}
