package main

import "errors"

type BlockValidationService struct {
	chain *Blockchain
}

func NewBlockValidator(chain *Blockchain) *BlockValidationService {
	return &BlockValidationService{chain: chain}
}

func (bv *BlockValidationService) ValidateFullChain() error {
	for i := 1; i < len(bv.chain.Chain); i++ {
		current := bv.chain.Chain[i]
		prev := bv.chain.Chain[i-1]

		if !isBlockValid(current, prev) {
			return errors.New("invalid block detected")
		}
		if current.Hash != calculateHash(current) {
			return errors.New("block hash mismatch")
		}
	}
	return nil
}

func (bv *BlockValidationService) ValidateSingleBlock(block Block) error {
	if block.Hash != calculateHash(block) {
		return errors.New("invalid block hash")
	}
	if !isValidHash(block.Hash, block.Difficulty) {
		return errors.New("invalid proof of work")
	}
	return nil
}
