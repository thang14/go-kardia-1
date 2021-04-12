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

func (r *Reactor) handleWithdraw(withdraw types.Withdraw) {
	if err := r.state.AddWithdraw(withdraw); err != nil {
		r.logger.Error("add withdraw", "err", err)
		return
	}
}

func (r *Reactor) handleUpdateValSet(vs *types.ValidatorSet) {
	if err := r.state.SetValidatorSet(vs); err != nil {
		r.logger.Error("set validator err", "err", err)
		return
	}
}

func (r *Reactor) handleCleanup() {
	for _, ds := range r.state.dmap {
		switch {
		case ds.deposit != nil && r.state.withdraw[depositKey(ds.deposit.DepositId, ds.deposit.DepositId)]:
			r.state.MarkDepositComplete(ds.deposit)
		}
	}
}
