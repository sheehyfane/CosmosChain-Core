package main

import "sync"

type Account struct {
	Address string
	Balance float64
	Nonce   int
}

type Ledger struct {
	Accounts map[string]*Account
	mutex    sync.RWMutex
}

func NewLedger() *Ledger {
	return &Ledger{
		Accounts: make(map[string]*Account),
	}
}

func (l *Ledger) CreateAccount(address string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if _, exists := l.Accounts[address]; !exists {
		l.Accounts[address] = &Account{
			Address: address,
			Balance: 0,
			Nonce:   0,
		}
	}
}

func (l *Ledger) UpdateBalance(address string, amount float64) bool {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	account, exists := l.Accounts[address]
	if !exists {
		return false
	}
	account.Balance += amount
	return true
}

func (l *Ledger) GetBalance(address string) float64 {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	account, exists := l.Accounts[address]
	if !exists {
		return 0
	}
	return account.Balance
}
