package main

import (
	"errors"
	"strings"
)

type Contract struct {
	ContractID string
	Code       string
	Owner      string
	IsActive   bool
}

type VM struct {
	Contracts map[string]*Contract
}

func NewVM() *VM {
	return &VM{
		Contracts: make(map[string]*Contract),
	}
}

func (vm *VM) DeployContract(contractID, code, owner string) {
	vm.Contracts[contractID] = &Contract{
		ContractID: contractID,
		Code:       code,
		Owner:      owner,
		IsActive:   true,
	}
}

func (vm *VM) ExecuteContract(contractID string, params []string) (string, error) {
	contract, exists := vm.Contracts[contractID]
	if !exists {
		return "", errors.New("contract not found")
	}
	if !contract.IsActive {
		return "", errors.New("contract disabled")
	}
	return "executed: " + strings.Join(params, ","), nil
}

func (vm *VM) DisableContract(contractID string) {
	if contract, exists := vm.Contracts[contractID]; exists {
		contract.IsActive = false
	}
}
