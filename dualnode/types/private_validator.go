package types

import (
	"crypto/ecdsa"

	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/crypto"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type PrivValidator interface {
	SignVote(vote *dproto.Vote) error
	GetAddress() common.Address
	GetPrivKey() *ecdsa.PrivateKey
}

type privValidator struct {
	privKey *ecdsa.PrivateKey
}

func (p *privValidator) SignVote(vote *dproto.Vote) error {
	return nil
}

// GetAddress ...
func (p *privValidator) GetAddress() common.Address {
	return crypto.PubkeyToAddress(p.GetPubKey())
}

func (p *privValidator) GetPrivKey() *ecdsa.PrivateKey {
	return p.privKey
}

func (p *privValidator) GetPubKey() ecdsa.PublicKey {
	return p.privKey.PublicKey
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

// GetAddress ...
func (p *MockPV) GetAddress() common.Address {
	return crypto.PubkeyToAddress(p.GetPubKey())
}

func (p *MockPV) GetPrivKey() *ecdsa.PrivateKey {
	return p.privKey
}

func (p *MockPV) GetPubKey() ecdsa.PublicKey {
	return p.privKey.PublicKey
}
