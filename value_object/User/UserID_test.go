package User

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyUserID(t *testing.T) {
	userID, err := NewUserID("")
	assert.NotNil(t, err)
	assert.Nil(t, userID)
}

func TestMultiByteCharacter(t *testing.T) {
	userID, err := NewUserID("あ")
	assert.NotNil(t, err)
	assert.Nil(t, userID)
}

func TestInvalidSymbol(t *testing.T) {
	userID, err := NewUserID("123#")
	assert.NotNil(t, err)
	assert.Nil(t, userID)
}

func TestValidUserID(t *testing.T) {
	userID, err := NewUserID("abc-ABC_123")
	assert.Nil(t, err)
	assert.NotNil(t, userID)
	assert.Equal(t, "abc-ABC_123", userID.GetValue())
}
