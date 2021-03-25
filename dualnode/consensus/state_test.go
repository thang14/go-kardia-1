package consensus

import (
	"testing"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	kardiachain_dualnode "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/stretchr/testify/assert"
)

func TestAddDeposit(t *testing.T) {
	pool := NewPool()
	s := store.New(memorydb.New())
	state, err := NewState(pool, s)
	assert.NoError(t, err)

	priv := types.NewMockPV()
	state.SetPrivValidator(priv)

	depositRecord := &kardiachain_dualnode.Deposit{
		Hash:        []byte("dfsafdsf"),
		Source:      1, // eth
		Destination: 2, // kai,
		DepositId:   3,
	}

	err = state.AddDeposit(depositRecord)
	assert.NoError(t, err)

	deposit := state.GetDepositByID(2, 3)
	assert.Equal(t, deposit.Hash, depositRecord.Hash)

	signs := state.GetDepositSignatures(deposit)
	assert.Equal(t, len(signs), 1)

	// mark deposit completed
	err = state.MarkDepositComplete(deposit)
	assert.NoError(t, err)

	signs = state.GetDepositSignatures(deposit)
	assert.Equal(t, len(signs), 0)
}
