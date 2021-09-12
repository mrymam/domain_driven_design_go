package Money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrencySame(t *testing.T) {
	jpen := NewCurrency(JpaneseEN)
	jpen2 := NewCurrency(JpaneseEN)

	assert.Equal(t, true, jpen.IsSame(jpen2))
	assert.Equal(t, true, jpen2.IsSame(jpen))
}

func TestCurrencyNotSame(t *testing.T) {
	jpen := NewCurrency(JpaneseEN)
	jpen2 := NewCurrency(Dollar)

	assert.Equal(t, false, jpen.IsSame(jpen2))
	assert.Equal(t, false, jpen2.IsSame(jpen))
}
