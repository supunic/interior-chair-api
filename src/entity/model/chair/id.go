package chair

import "errors"

type ID uint

func NewChairID(val uint) (*ID, error) {
	if val < 1 {
		return nil, errors.New("idは1以上の整数である必要があります。")
	}
	id := ID(val)
	return &id, nil
}

func (v *ID) Value() uint {
	return uint(*v)
}
