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
			r.handleDeposit(depositRecord)
		case valSet := <-r.valSetC:
			r.handleUpdateValSet(valSet)
		case withdraw := <-r.withdrawC:
			r.handleWithdraw(withdraw)
		case <-cleanup.C:
			r.handleCleanup()
		case <-r.Quit():
			return nil
		}
	}
}

func (r *Reactor) handleDeposit(d *dproto.Deposit) {
	if err := r.state.AddDeposit(d); err != nil {
		r.logger.Error("add deposit", "err", err)
		return
	}
}

func (r *Reactor) handleWithdraw(withdraw types.Withdraw) error {
	deposit := r.state.GetDepositByID(withdraw.DestChainId, withdraw.DepositId)
	if deposit == nil {
		return nil
	}
	return r.state.MarkDepositComplete(deposit)
}

func (r *Reactor) handleUpdateValSet(vs *types.ValidatorSet) {
	if err := r.state.SetValidatorSet(vs); err != nil {
		r.logger.Error("set validator err", "err", err)
		return
	}
}
