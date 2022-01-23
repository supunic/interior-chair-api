package chairAuthor

type Image string

func NewChairAuthorImage(val string) (*Image, error) {
	image := Image(val)
	return &image, nil
}

func (v *Image) Value() string {
	return string(*v)
}
