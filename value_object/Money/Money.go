package Money

import "fmt"

type Money struct {
	amount   uint64
	currency *Currency
}

func NewMoney(amount uint64, currency *Currency) *Money {
	return &Money{
		amount:   amount,
		currency: currency,
	}
}

func (m Money) GetAmount() uint64 {
	return m.amount
}

func (m Money) GetCurrencyName() string {
	return m.currency.GetCurrencyName()
}

func (m Money) Add(other *Money) (*Money, error) {
	if !m.currency.IsSame(other.currency) {
		return nil, fmt.Errorf("not same currency")
	}

	money := NewMoney(m.amount+other.amount, m.currency)
	return money, nil
}
