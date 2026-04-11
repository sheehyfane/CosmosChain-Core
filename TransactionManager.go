package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Transaction struct {
	TxID        string
	Sender      string
	Receiver    string
	Amount      float64
	Fee         float64
	Timestamp   string
	Signature   string
	IsConfirmed bool
}

type TxManager struct {
	PendingTxs []Transaction
	Confirmed  []Transaction
}

func (tm *TxManager) CreateTransaction(sender, receiver string, amount, fee float64) Transaction {
	tx := Transaction{
		Sender:    sender,
		Receiver:  receiver,
		Amount:    amount,
		Fee:       fee,
		Timestamp: time.Now().String(),
	}
	hash := sha256.Sum256([]byte(sender + receiver + string(amount) + tx.Timestamp))
	tx.TxID = hex.EncodeToString(hash[:])
	return tx
}

func (tm *TxManager) AddPendingTx(tx Transaction) {
	tm.PendingTxs = append(tm.PendingTxs, tx)
}

func (tm *TxManager) ConfirmTx(txID string) {
	for i, tx := range tm.PendingTxs {
		if tx.TxID == txID {
			tm.PendingTxs[i].IsConfirmed = true
			tm.Confirmed = append(tm.Confirmed, tm.PendingTxs[i])
			tm.PendingTxs = append(tm.PendingTxs[:i], tm.PendingTxs[i+1:]...)
			break
		}
	}
}
