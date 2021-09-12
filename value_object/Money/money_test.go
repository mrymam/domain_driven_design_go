package Money

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	jpen := NewCurrency(JpaneseEN)

	money := NewMoney(100, jpen)
	money2 := NewMoney(200, jpen)
	money, _ = money.Add(money2)

	assert.Equal(t, uint64(300), money.GetAmount())
}

func TestAddFail(t *testing.T) {
	jpen := NewCurrency(JpaneseEN)
	dollar := NewCurrency(Dollar)

	money := NewMoney(100, jpen)
	money2 := NewMoney(100, dollar)
	_, err := money.Add(money2)

	assert.NotNil(t, err)
}
