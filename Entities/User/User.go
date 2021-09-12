package User

import (
	UserValues "github.com/onyanko-pon/domain_driven_design_go/value_object/User"
)

type User struct {
	id   UserValues.UserID
	name UserValues.UserName
}

func NewUser(idStr string, nameStr string) (*User, error) {

	id, err := UserValues.NewUserID(idStr)
	if err != nil {
		return nil, err
	}

	name, err := UserValues.NewUserName(nameStr)
	if err != nil {
		return nil, err
	}

	return &User{
		id:   *id,
		name: *name,
	}, nil
}

func (u *User) UpdateName(nameStr string) error {
	newName, err := UserValues.NewUserName(nameStr)
	if err != nil {
		return err
	}
	u.name = *newName
	return nil
}

func (u User) GetName() UserValues.UserName {
	return u.name
}

func (u User) GetID() UserValues.UserID {
	return u.id
}

func (u User) Equal(other User) bool {
	return u.id.Equal(other.GetID())
}
