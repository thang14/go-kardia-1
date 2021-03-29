package types

import "github.com/kardiachain/go-kardia/lib/common"

type ValidatorSet struct {
	validators []common.Address
}

func (vs *ValidatorSet) Add(address common.Address) {
	if !vs.Has(address) {
		vs.validators = append(vs.validators, address)
	}
}

func (vs *ValidatorSet) Remove(addr common.Address) {
	newVals := make([]common.Address, 0)
	for _, val := range vs.validators {
		if !val.Equal(addr) {
			newVals = append(newVals, val)
		}
	}
	vs.validators = newVals
}

func (vs *ValidatorSet) Has(addr common.Address) bool {
	for _, val := range vs.validators {
		if val.Equal(addr) {
			return true
		}
	}
	return false
}

func NewValidatorSet(addrs []common.Address) *ValidatorSet {
	return &ValidatorSet{validators: addrs}
}
