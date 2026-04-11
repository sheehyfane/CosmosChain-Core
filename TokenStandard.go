package main

import "errors"

type Token struct {
	Name     string
	Symbol   string
	Decimals int
	Supply   float64
	Balances map[string]float64
}

func NewToken(name, symbol string, decimals int, supply float64) *Token {
	return &Token{
		Name:     name,
		Symbol:   symbol,
		Decimals: decimals,
		Supply:   supply,
		Balances: map[string]float64{},
	}
}

func (t *Token) Transfer(from, to string, amount float64) error {
	if t.Balances[from] < amount {
		return errors.New("insufficient balance")
	}
	t.Balances[from] -= amount
	t.Balances[to] += amount
	return nil
}

func (t *Token) Mint(address string, amount float64) {
	t.Balances[address] += amount
	t.Supply += amount
}
