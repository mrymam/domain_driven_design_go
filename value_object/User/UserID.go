package User

import (
	"fmt"
	"regexp"
)

type UserID struct {
	value string
}

func NewUserID(value string) (*UserID, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf("empty UserID: %s", value)
	}
	// 数字、アルファベット、- _で構成されているかどうか
	re := regexp.MustCompile(`^[a-zA-Z0-9\-\_]+$`)
	res := re.MatchString(value)

	if !res {
		return nil, fmt.Errorf("invalid characters in %s. UserID must be consisted by 0-9, a-z, A-Z, \"-\" and \"_\"", value)
	}

	return &UserID{
		value: value,
	}, nil
}

func (u UserID) GetValue() string {
	return u.value
}
