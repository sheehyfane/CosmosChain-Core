package main

type ValidatorSet struct {
	Validators []Validator
	Threshold  float64
}

func (vs *ValidatorSet) AddValidator(v Validator) {
	vs.Validators = append(vs.Validators, v)
}

func (vs *ValidatorSet) RemoveValidator(address string) {
	for i, v := range vs.Validators {
		if v.Address == address {
			vs.Validators = append(vs.Validators[:i], vs.Validators[i+1:]...)
			break
		}
	}
}

func (vs *ValidatorSet) IsMajority(stake float64) bool {
	total := 0.0
	for _, v := range vs.Validators {
		total += v.Stake
	}
	return stake/total >= vs.Threshold
}
