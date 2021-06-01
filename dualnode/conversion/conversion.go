package conversion

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"

	btss "github.com/binance-chain/tss-lib/tss"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func GetParties(keys []string, localPartyKey string) ([]*btss.PartyID, *btss.PartyID, error) {
	var localPartyID *btss.PartyID
	var unSortedPartiesID []*btss.PartyID
	sort.Strings(keys)
	for idx, item := range keys {
		pk, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, item)
		if err != nil {
			return nil, nil, fmt.Errorf("fail to get account pub key address(%s): %w", item, err)
		}
		key := new(big.Int).SetBytes(pk.Bytes())
		// Set up the parameters
		// Note: The `id` and `moniker` fields are for convenience to allow you to easily track participants.
		// The `id` should be a unique string representing this party in the network and `moniker` can be anything (even left blank).
		// The `uniqueKey` is a unique identifying key for this peer (such as its p2p public key) as a big.Int.
		partyID := btss.NewPartyID(strconv.Itoa(idx), "", key)
		if item == localPartyKey {
			localPartyID = partyID
		}
		unSortedPartiesID = append(unSortedPartiesID, partyID)
	}
	if localPartyID == nil {
		return nil, nil, errors.New("local party is not in the list")
	}

	partiesID := btss.SortPartyIDs(unSortedPartiesID)
	return partiesID, localPartyID, nil
}

func BytesKeysToString(rawArr [][]byte, raw []byte) ([]string, string) {
	resultArr := make([]string, len(rawArr))
	for i := range rawArr {
		resultArr[i] = string(rawArr[i])
	}
	result := string(raw)
	return resultArr, result
}
