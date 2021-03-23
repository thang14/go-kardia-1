package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/store"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

type API struct {
	store *store.Store
}

func (api *API) GetDeposit(hash []byte) (*dproto.Deposit, error) {
	return api.store.GetPendingDeposit(hash)
}
