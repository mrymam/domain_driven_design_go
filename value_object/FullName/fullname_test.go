package FullName

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFullName(t *testing.T) {
	fullName := NewFullName("Tsukasa", "Maruyama")
	assert.Equal(t, "Maruyama Tsukasa", fullName.GetFullName())
}
