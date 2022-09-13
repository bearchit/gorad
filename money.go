package gorad

import (
	"math/big"

	"golang.org/x/text/currency"
)

type Money struct {
	Amount   *big.Int
	Currency currency.Unit
}

func NewMoney(amount int64, currency currency.Unit) Money {
	return Money{
		Amount:   big.NewInt(amount),
		Currency: currency,
	}
}
