package chair

import "errors"

type ID int

func NewChairID(val int) (*ID, error) {
	if val < 1 {
		return nil, errors.New("idは1以上の整数である必要があります。")
	}
	id := ID(val)
	return &id, nil
}
