package author

type Image string

func NewAuthorImage(val string) (*Image, error) {
	image := Image(val)
	return &image, nil
}

func (v *Image) Value() string {
	return string(*v)
}
