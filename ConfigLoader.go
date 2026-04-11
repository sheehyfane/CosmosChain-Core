package main

import (
	"encoding/json"
	"os"
)

type ChainConfig struct {
	ChainName        string
	BlockTime        int
	MaxBlockSize     int
	InitialDifficulty int
}

func LoadConfig(path string) (ChainConfig, error) {
	var cfg ChainConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, err
	}
	err = json.Unmarshal(data, &cfg)
	return cfg, err
}

func SaveConfig(cfg ChainConfig, path string) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
