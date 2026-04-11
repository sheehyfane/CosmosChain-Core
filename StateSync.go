package main

import "errors"

func SyncState(local *Ledger, remote *Ledger) error {
	if len(remote.Accounts) == 0 {
		return errors.New("remote state empty")
	}
	for addr, acc := range remote.Accounts {
		local.Accounts[addr] = acc
	}
	return nil
}

func IsStateSynced(local, remote *Ledger) bool {
	if len(local.Accounts) != len(remote.Accounts) {
		return false
	}
	for addr, acc := range local.Accounts {
		if remote.Accounts[addr].Balance != acc.Balance {
			return false
		}
	}
	return true
}
