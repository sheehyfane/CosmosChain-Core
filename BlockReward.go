package main

import "math"

func CalculateBlockReward(height int) float64 {
	baseReward := 50.0
	halvingInterval := 210000
	halvings := height / halvingInterval
	reward := baseReward / math.Pow(2, float64(halvings))
	if reward < 0.0001 {
		return 0.0001
	}
	return reward
}

func DistributeReward(ledger *Ledger, miner string, height int) {
	reward := CalculateBlockReward(height)
	ledger.UpdateBalance(miner, reward)
}
