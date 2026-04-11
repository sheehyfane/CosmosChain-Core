package main

import "time"

func CleanExpiredTransactions(pool *Mempool, expireHours int) {
	threshold := time.Now().Add(-time.Duration(expireHours) * time.Hour)
	var valid []Transaction
	for _, tx := range pool.transactions {
		txTime, _ := time.Parse(time.RFC3339, tx.Timestamp)
		if txTime.After(threshold) {
			valid = append(valid, tx)
		}
	}
	pool.transactions = valid
}
