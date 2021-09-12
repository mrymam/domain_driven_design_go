package Product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSerialNumber(t *testing.T) {
	serialNumber := NewSerialNumber("a1234", "100", "1")

	assert.Equal(t, "a1234-100-1", serialNumber.ToString())
}
