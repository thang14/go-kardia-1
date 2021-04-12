package consensus

import (
	"testing"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	"github.com/kardiachain/go-kardia/lib/common"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlers(t *testing.T) {
	s := store.New(memorydb.New())
	pool := NewPool()
	priv := types.NewMockPV()
	state, err := NewState(pool, s)
	require.NoError(t, err)
	state.SetPrivValidator(priv)

	depositC := make(chan *dproto.Deposit, 0)
	withdrawC := make(chan types.Withdraw, 0)
	vsC := make(chan *types.ValidatorSet, 0)

	r := newReactor(state, nil, depositC, withdrawC, vsC)

	// add deposit
	deposit := &dproto.Deposit{
		SourceChainId: 1,
		DestChainId:   1,
		DepositId:     1,
	}
	types.WithDepositHash(deposit, common.Address{})
	depositHash := common.BytesToHash(deposit.Hash)
	r.handleDeposit(deposit)
	ds := r.state.dmap[depositHash]
	assert.Equal(t, ds.deposit, deposit)

	// withdraw
	widthdraw := types.Withdraw{SourceChainId: deposit.SourceChainId, DestChainId: deposit.DestChainId, DepositId: deposit.DepositId}
	r.handleWithdraw(widthdraw)

	ds = r.state.dmap[depositHash]
	assert.Empty(t, ds)

	// update validator set
	priv2 := types.NewMockPV()
	r.handleUpdateValSet(types.NewValidatorSet([]common.Address{priv.GetAddress(), priv2.GetAddress()}))
	assert.True(t, r.state.vs.Has(priv.GetAddress()))
	assert.True(t, r.state.vs.Has(priv2.GetAddress()))
}
