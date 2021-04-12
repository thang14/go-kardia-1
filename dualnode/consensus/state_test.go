package consensus

import (
	"testing"

	"github.com/kardiachain/go-kardia/dualnode/store"
	"github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/kai/kaidb/memorydb"
	"github.com/kardiachain/go-kardia/lib/common"
	kardiachain_dualnode "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddDeposit(t *testing.T) {
	pool := NewPool()
	s := store.New(memorydb.New())
	state, err := NewState(pool, s)
	assert.NoError(t, err)

	priv := types.NewMockPV()
	priv2 := types.NewMockPV()
	state.SetPrivValidator(priv)
	state.SetValidatorSet(types.NewValidatorSet([]common.Address{priv.GetAddress()}))

	depositRecord := &kardiachain_dualnode.Deposit{
		SourceChainId: 1, // eth
		DestChainId:   2, // kai,
		DepositId:     3,
	}

	err = types.WithDepositHash(depositRecord, common.Address{})
	require.NoError(t, err)

	err = state.AddDeposit(depositRecord)
	assert.NoError(t, err)

	deposit := state.GetDepositByID(2, 3)
	assert.Equal(t, deposit.Hash, depositRecord.Hash)

	signs := state.Signs(deposit)
	assert.Equal(t, len(signs), 1)

	// add other vote
	vote := &kardiachain_dualnode.Vote{
		Hash: deposit.Hash,
		Addr: []byte("0x2"),
	}
	state.addVote(vote)
	// no add duplicated vote
	state.addVote(vote)

	vote2 := &kardiachain_dualnode.Vote{
		Hash: deposit.Hash,
		Addr: priv.GetAddress().Bytes(),
	}
	err = priv2.SignVote(vote2)
	assert.NoError(t, err)

	// invalid signature
	err = state.AddVote(vote2)
	assert.Error(t, err, "invalid signature")

	signs = state.Signs(deposit)
	assert.Equal(t, len(signs), 2)

	// mark deposit completed
	err = state.MarkDepositComplete(deposit)
	assert.NoError(t, err)

	signs = state.Signs(deposit)
	assert.Equal(t, len(signs), 0)
}
