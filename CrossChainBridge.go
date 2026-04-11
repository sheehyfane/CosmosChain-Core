package main

import "errors"

type CrossChainTx struct {
	SourceChain      string
	TargetChain      string
	SourceTxID       string
	TargetTxID       string
	Amount           float64
	Status           string
}

type Bridge struct {
	ChainIDs    []string
	PendingCross []CrossChainTx
}

func (b *Bridge) RegisterChain(chainID string) {
	b.ChainIDs = append(b.ChainIDs, chainID)
}

func (b *Bridge) InitiateCrossTransfer(source, target, sourceTx string, amount float64) string {
	crossTx := CrossChainTx{
		SourceChain: source,
		TargetChain: target,
		SourceTxID:  sourceTx,
		Amount:      amount,
		Status:      "pending",
	}
	b.PendingCross = append(b.PendingCross, crossTx)
	return "cross_chain_initiated"
}

func (b *Bridge) CompleteCrossTransfer(sourceTx, targetTx string) error {
	for i, tx := range b.PendingCross {
		if tx.SourceTxID == sourceTx {
			b.PendingCross[i].TargetTxID = targetTx
			b.PendingCross[i].Status = "completed"
			return nil
		}
	}
	return errors.New("transaction not found")
}
