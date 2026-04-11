package main

import (
	"encoding/json"
	"os"
)

type Storage struct {
	FilePath string
}

func (s *Storage) SaveBlock(block Block) error {
	data, err := json.Marshal(block)
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath+"/block_"+string(block.Index)+".json", data, 0644)
}

func (s *Storage) LoadBlock(index int) (Block, error) {
	var block Block
	data, err := os.ReadFile(s.FilePath + "/block_" + string(index) + ".json")
	if err != nil {
		return block, err
	}
	err = json.Unmarshal(data, &block)
	return block, err
}

func (s *Storage) SaveState(ledger Ledger) error {
	data, err := json.Marshal(ledger)
	if err != nil {
		return err
	}
	return os.WriteFile(s.FilePath+"/state.json", data, 0644)
}
