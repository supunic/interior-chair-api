package gameType

type GameType struct {
	ID   ID   `json:"id"`
	Name Name `json:"name"`
}

func NewGameType(id int, name string) (*GameType, error) {
	gameTypeId, err := NewGameTypeID(id)
	if err != nil {
		return nil, err
	}

	gameTypeName, err := NewGameTypeName(name)
	if err != nil {
		return nil, err
	}

	return &GameType{
		ID:   *gameTypeId,
		Name: *gameTypeName,
	}, nil
}
