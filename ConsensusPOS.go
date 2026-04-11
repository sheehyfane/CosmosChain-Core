package main

import (
	"math/rand"
	"time"
)

type Validator struct {
	Address    string
	Stake      float64
	IsActive   bool
}

type POSConsensus struct {
	Validators []Validator
}

func (pos *POSConsensus) RegisterValidator(address string, stake float64) {
	pos.Validators = append(pos.Validators, Validator{
		Address:  address,
		Stake:    stake,
		IsActive: true,
	})
}

func (pos *POSConsensus) SelectBlockProducer() string {
	var activeValidators []Validator
	for _, v := range pos.Validators {
		if v.IsActive && v.Stake > 0 {
			activeValidators = append(activeValidators, v)
		}
	}

	if len(activeValidators) == 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())
	totalStake := 0.0
	for _, v := range activeValidators {
		totalStake += v.Stake
	}

	r := rand.Float64() * totalStake
	current := 0.0
	for _, v := range activeValidators {
		current += v.Stake
		if current >= r {
			return v.Address
		}
	}
	return activeValidators[0].Address
}

func (pos *POSConsensus) SlashValidator(address string, penalty float64) {
	for i := range pos.Validators {
		if pos.Validators[i].Address == address {
			pos.Validators[i].Stake -= penalty
			if pos.Validators[i].Stake < 0 {
				pos.Validators[i].IsActive = false
			}
			break
		}
	}
}
