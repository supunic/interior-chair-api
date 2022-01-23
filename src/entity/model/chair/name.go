package chair

import "errors"

type Name string

func NewChairName(val string) (*Name, error) {
	if len(val) < 2 {
		return nil, errors.New("nameは3文字以上の文字列でなければなりません。")
	}
	name := Name(val)
	return &name, nil
}

func (v *Name) Value() string {
	return string(*v)
}
