package main

import (
	"fmt"

	"github.com/onyanko-pon/domain_driven_design_go/value_object/FullName"
	"github.com/onyanko-pon/domain_driven_design_go/value_object/Money"
)

func main() {
	fullName := FullName.NewFullName("Tsukasa", "Maruyama")
	fmt.Printf("The Name is %s.\n", fullName.GetFullName())

	money := Money.NewMoney(100, Money.NewCurrency(Money.JpaneseEN))
	fmt.Printf("Money amount: %d, currency: %s.\n", money.GetAmount(), money.GetCurrencyName())
	money2 := Money.NewMoney(200, Money.NewCurrency(Money.JpaneseEN))

	money, _ = money.Add(money2)
	fmt.Printf("Added money amount: %d, currency: %s.\n", money.GetAmount(), money.GetCurrencyName())
}
