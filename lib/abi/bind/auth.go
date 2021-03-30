/*
 *  Copyright 2018 KardiaChain
 *  This file is part of the go-kardia library.
 *
 *  The go-kardia library is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Lesser General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  The go-kardia library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU Lesser General Public License for more details.
 *
 *  You should have received a copy of the GNU Lesser General Public License
 *  along with the go-kardia library. If not, see <http://www.gnu.org/licenses/>.
 */

package bind

import (
	"crypto/ecdsa"
	"errors"

	"github.com/kardiachain/go-kardia/lib/common"
	"github.com/kardiachain/go-kardia/lib/crypto"
	"github.com/kardiachain/go-kardia/lib/log"
	"github.com/kardiachain/go-kardia/types"
)

// ErrNotAuthorized is returned when an account is not properly unlocked.
var ErrNotAuthorized = errors.New("not authorized to sign this account")

// NewKeyedTransactor is a utility method to easily create a transaction signer
// from a single private key.
//
// Deprecated: Use NewKeyedTransactorWithChainID instead.
func NewKeyedTransactor(key *ecdsa.PrivateKey) *TransactOpts {
	log.Warn("WARNING: NewKeyedTransactor has been deprecated in favour of NewKeyedTransactorWithChainID")
	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
	signer := types.HomesteadSigner{}
	return &TransactOpts{
		From: keyAddr,
		Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
			if address != keyAddr {
				return nil, ErrNotAuthorized
			}
			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
			if err != nil {
				return nil, err
			}
			return tx.WithSignature(signer, signature)
		},
	}
}
