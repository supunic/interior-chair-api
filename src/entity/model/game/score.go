package game

type Score int

func NewGameScore(val int) (*Score, error) {
	score := Score(val)
	return &score, nil
}
