package User

import "fmt"

const MIN_USERNAME_LENGTH = 3

type UserName struct {
	value string
}

func NewUserName(value string) (*UserName, error) {
	if len(value) < MIN_USERNAME_LENGTH {
		return nil, fmt.Errorf("too short UserName: %s. must be longer than %d", value, MIN_USERNAME_LENGTH)
	}
	return &UserName{
		value: value,
	}, nil
}

func (u UserName) GetValue() string {
	return u.value
}
