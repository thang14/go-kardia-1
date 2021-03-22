package dualnode

import (
	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
)

type API struct {
	store *store.Store
}

func (api *API) GetDeposit(hash string) (*types.Deposit, error) {
	return api.store.GetDeposit(hash)
}
