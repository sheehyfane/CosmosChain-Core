package main

import "errors"

type UTXO struct {
	TxID      string
	Index     int
	Amount    float64
	Spent     bool
	Owner     string
}

type UTXOSet struct {
	UTXOs []UTXO
}

func (u *UTXOSet) AddUTXO(txID string, index int, amount float64, owner string) {
	u.UTXOs = append(u.UTXOs, UTXO{
		TxID:   txID,
		Index:  index,
		Amount: amount,
		Owner:  owner,
		Spent:  false,
	})
}

func (u *UTXOSet) SpendUTXO(txID string, index int) error {
	for i := range u.UTXOs {
		if u.UTXOs[i].TxID == txID && u.UTXOs[i].Index == index && !u.UTXOs[i].Spent {
			u.UTXOs[i].Spent = true
			return nil
		}
	}
	return errors.New("utxo not found")
}
