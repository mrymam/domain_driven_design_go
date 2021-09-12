package User

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortUserName(t *testing.T) {
	userName, err := NewUserName("me")
	assert.NotNil(t, err)
	assert.Nil(t, userName)
}

func TestSuccessUserName(t *testing.T) {
	userName, err := NewUserName("123")
	assert.Nil(t, err)
	assert.NotNil(t, userName)
	assert.Equal(t, "123", userName.GetValue())
}
