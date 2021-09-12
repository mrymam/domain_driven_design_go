package User

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidUser(t *testing.T) {
	user, err := NewUser("userid", "username")

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestNewUserFailInvalidUserName(t *testing.T) {
	user, err := NewUser("userid", "")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestNewUserFailInvalidUserID(t *testing.T) {
	user, err := NewUser("インバリッド", "username")

	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestUpdateUserNameSuccess(t *testing.T) {
	user, _ := NewUser("userid", "username")
	err := user.UpdateName("new-username")

	assert.Nil(t, err)
	assert.Equal(t, "new-username", user.GetName().GetValue())
}

func TestUpdateUserNameFail(t *testing.T) {
	user, _ := NewUser("userid", "username")
	err := user.UpdateName("")

	assert.NotNil(t, err)
	assert.Equal(t, "username", user.GetName().GetValue())
}

func TestEqualUser(t *testing.T) {
	user, _ := NewUser("userid", "username")
	user2, _ := NewUser("userid", "username2")

	assert.True(t, user.Equal(*user2))
}
