package author

import "errors"

type ID uint

func NewAuthorID(val uint) (*ID, error) {
	if val < 0 {
		return nil, errors.New("idは0以上の整数である必要があります。")
	}
	id := ID(val)
	return &id, nil
}

func (v *ID) Value() uint {
	return uint(*v)
}
