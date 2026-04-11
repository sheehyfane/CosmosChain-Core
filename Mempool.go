package main

import (
	"sort"
	"sync"
)

type Mempool struct {
	transactions []Transaction
	mutex        sync.RWMutex
}

func NewMempool() *Mempool {
	return &Mempool{
		transactions: []Transaction{},
	}
}

func (mp *Mempool) AddTransaction(tx Transaction) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	mp.transactions = append(mp.transactions, tx)
}

func (mp *Mempool) GetTopTransactions(limit int) []Transaction {
	mp.mutex.RLock()
	defer mp.mutex.RUnlock()
	sort.Slice(mp.transactions, func(i, j int) bool {
		return mp.transactions[i].Fee > mp.transactions[j].Fee
	})
	if limit > len(mp.transactions) {
		limit = len(mp.transactions)
	}
	return mp.transactions[:limit]
}

func (mp *Mempool) RemoveTransactions(txs []Transaction) {
	mp.mutex.Lock()
	defer mp.mutex.Unlock()
	temp := []Transaction{}
	for _, tx := range mp.transactions {
		found := false
		for _, t := range txs {
			if tx.TxID == t.TxID {
				found = true
				break
			}
		}
		if !found {
			temp = append(temp, tx)
		}
	}
	mp.transactions = temp
}
