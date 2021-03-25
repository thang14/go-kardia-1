package types

import (
	"crypto/ecdsa"

	"github.com/kardiachain/go-kardia/lib/crypto"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type PrivValidator interface {
	SignVote(vote *dproto.Vote) error
}

type privValidator struct {
}

func (p *privValidator) SignVote(vote *dproto.Vote) error {
	return nil
}

func NewPrivValidator() PrivValidator {
	return &privValidator{}
}

//----------------------------------------
// MockPV

// MockPV implements PrivValidator without any safety or persistence.
// Only use it for testing.
type MockPV struct {
	privKey *ecdsa.PrivateKey
}

// NewMockPV new mock priv validator
func NewMockPV() *MockPV {
	priv, err := crypto.GenerateKey()

	if err != nil {
		panic(err)
	}

	return &MockPV{priv}
}

func (p *MockPV) SignVote(vote *dproto.Vote) error {
	return nil
}
