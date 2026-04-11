package main

import "math"

type ChainStats struct {
	TotalTx    int
	AvgBlockSize float64
	TotalSupply float64
}

func (cs *ChainStats) ComputeStats(chain []Block, txs []Transaction) {
	cs.TotalTx = len(txs)
	sum := 0
	for _, b := range chain {
		sum += len(b.Data)
	}
	cs.AvgBlockSize = math.Round(float64(sum)/float64(len(chain))*100) / 100
}
