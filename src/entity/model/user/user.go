package user

type User struct {
	ID   ID   `json:"id"`
	Name Name `json:"name"`
}

func NewUser(id int, name string) (*User, error) {
	UserId, err := NewUserID(id)
	if err != nil {
		return nil, err
	}

	UserName, err := NewUserName(name)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:   *UserId,
		Name: *UserName,
	}, nil
}
