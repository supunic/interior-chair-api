package chair

type Image string

func NewChairImage(val string) (*Image, error) {
	image := Image(val)
	return &image, nil
}

func (v *Image) Value() string {
	return string(*v)
}
