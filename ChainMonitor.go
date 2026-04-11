package main

import "fmt"

type ChainMonitor struct {
	chain *Blockchain
}

func NewChainMonitor(chain *Blockchain) *ChainMonitor {
	return &ChainMonitor{chain: chain}
}

func (cm *ChainMonitor) GetChainHeight() int {
	return len(cm.chain.Chain) - 1
}

func (cm *ChainMonitor) PrintChainInfo() {
	fmt.Printf("chain height: %d\n", cm.GetChainHeight())
	fmt.Printf("current difficulty: %d\n", cm.chain.Difficulty)
}

func (cm *ChainMonitor) CheckFork() bool {
	for i := 1; i < len(cm.chain.Chain); i++ {
		if cm.chain.Chain[i].PrevHash != cm.chain.Chain[i-1].Hash {
			return true
		}
	}
	return false
}
