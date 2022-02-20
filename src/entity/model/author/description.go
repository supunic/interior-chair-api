package author

type Description string

func NewAuthorDescription(val string) (*Description, error) {
	description := Description(val)
	return &description, nil
}

func (v *Description) Value() string {
	return string(*v)
}
