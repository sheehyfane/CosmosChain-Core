package main

import (
	"time"
)

func CreateGenesisBlock() Block {
	genesis := Block{
		Index:      0,
		Timestamp:  time.Now().String(),
		Data:       "genesis block - chain initialization",
		PrevHash:   "0",
		Difficulty: 2,
	}
	genesis.Hash = calculateHash(genesis)
	return genesis
}

func InitBlockchain() *Blockchain {
	genesis := CreateGenesisBlock()
	return &Blockchain{
		Chain:      []Block{genesis},
		Difficulty: 2,
	}
}
