package User

type User struct {
	id   UserID
	name UserName
}

func NewUser(id UserID, name UserName) (*User, error) {
	return &User{
		id:   id,
		name: name,
	}, nil
}
