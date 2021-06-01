package tss

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/tss"

	tssAbi "github.com/kardiachain/go-kardia/dualnode/chains/kardiachain/tss"
	"github.com/kardiachain/go-kardia/dualnode/conversion"
	"github.com/kardiachain/go-kardia/lib/crypto"
)

func (r *Reactor) keygen(vaultEv *tssAbi.TssVaultCreated) error {
	var (
		partyID    *tss.PartyID
		privPubKey = r.privValidator.GetPubKey()
	)
	unsortedPIDs := make([]*tss.PartyID, len(vaultEv.Members))
	for i, member := range vaultEv.Members {
		pID := tss.NewPartyID(vaultEv.Id.String(), string(vaultEv.PubKey), new(big.Int).SetBytes(member))
		unsortedPIDs[i] = pID
		if bytes.Equal(member, crypto.FromECDSAPub(&privPubKey)) { // determine the partyID if this private validator pubkey is in members list
			partyID = pID

		}
	}
	partiesID, localPartyID, err := conversion.GetParties(conversion.BytesKeysToString(vaultEv.Members, privPubKey.))
	if err != nil {
		return nil, fmt.Errorf("fail to get keygen parties: %w", err)
	}
	pIDs := tss.SortPartyIDs(unsortedPIDs)
	p2pCtx := tss.NewPeerContext(pIDs)
	threshold := len(vaultEv.Members)/2 + 1
	params := tss.NewParameters(p2pCtx, partyID, len(pIDs), threshold)
	fixtures, pIDs, err := LoadKeygenTestFixtures(testParticipants)
	if err != nil {
		common.Logger.Info("No test fixtures were found, so the safe primes will be generated from scratch. This may take a while...")
		pIDs = tss.GenerateTestPartyIDs(testParticipants)
	}
}
