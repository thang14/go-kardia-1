package tss

import (
	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/binance-chain/tss-lib/tss"
)

type State struct {
	height             int64
	partyIDs           tss.SortedPartyIDs
	localPartyID       *tss.PartyID
	localPartySaveData keygen.LocalPartySaveData
}
