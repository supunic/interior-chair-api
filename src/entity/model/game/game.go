package game

type Game struct {
	ID    ID    `json:"id"`
	Score Score `json:"score"`
}

func NewGame(id int, score int) (*Game, error) {
	gameId, err := NewGameID(id)
	if err != nil {
		return nil, err
	}

	gameScore, err := NewGameScore(score)
	if err != nil {
		return nil, err
	}

	return &Game{
		ID:    *gameId,
		Score: *gameScore,
	}, nil
}
