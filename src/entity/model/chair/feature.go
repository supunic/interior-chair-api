package chair

type Feature string

func NewChairFeature(val string) (*Feature, error) {
	feature := Feature(val)
	return &feature, nil
}

func (v *Feature) Value() string {
	return string(*v)
}
