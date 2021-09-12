package Money

type CurrencyType int

const (
	JpaneseEN CurrencyType = iota
	Dollar
)

func NewCurrency(currencyType CurrencyType) *Currency {
	return &Currency{
		CurrencyType: currencyType,
	}
}

type Currency struct {
	CurrencyType CurrencyType
}

func (c Currency) GetCurrencyName() string {
	switch c.CurrencyType {
	case JpaneseEN:
		return "円"
	case Dollar:
		return "ドル"
	}
	return ""
}

func (c Currency) IsSame(other *Currency) bool {
	return c.CurrencyType == other.CurrencyType
}
