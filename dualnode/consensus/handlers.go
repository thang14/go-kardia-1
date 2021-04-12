package consensus

import (
	"time"

	"github.com/kardiachain/go-kardia/dualnode/types"
	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

func (r *Reactor) run() error {
	cleanup := time.NewTicker(30 * time.Second)
	for {
		select {
		case depositRecord := <-r.depositC:
			return r.handleDeposit(depositRecord)
		case valSet := <-r.valSetC:
			return r.handleUpdateValSet(valSet)
		case withdraw := <-r.withdrawC:
			return r.handleWithdraw(withdraw)
		case <-cleanup.C:
			return r.handleCleanup()
		case <-r.Quit():
			return nil
		}
	}
}

func (r *Reactor) handleDeposit(d *dproto.Deposit) error {
	return r.state.AddDeposit(d)
}

func (r *Reactor) handleWithdraw(withdraw types.Withdraw) error {
	deposit := r.state.GetDepositByID(withdraw.DestChainId, withdraw.DepositId)
	return r.state.MarkDepositComplete(deposit)
}

func (r *Reactor) handleUpdateValSet(vs *types.ValidatorSet) error {
	return r.state.SetValidatorSet(vs)
}
