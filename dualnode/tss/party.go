package tss

import (
	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	"github.com/binance-chain/tss-lib/tss"

	kcommon "github.com/kardiachain/go-kardia/lib/common"
)

func (r *Reactor) handlerSigning(msg []byte) {
	// r.router.Send(msg, signatures);
	peerCtx := tss.NewPeerContext(r.state.partyIDs)
	totalParties := len(r.state.partyIDs)
	params := tss.NewParameters(peerCtx, r.state.localPartyID, totalParties, totalParties/2+1)

	outCh := make(chan tss.Message)
	endCh := make(chan common.SignatureData)
	errCh := make(chan *tss.Error)

	party := signing.NewLocalParty(kcommon.HashToInt(msg), params, r.state.localPartySaveData, outCh, endCh)

	r.signingPartyM.Lock()
	r.signingParties[kcommon.BytesToHash(msg)] = party
	r.signingPartyM.Unlock()
	for {
		select {
		case err := <-errCh:
			r.logger.Error("handle signing", "err", err)
			return
		case <-outCh:
			// add message to pool
			return
		case sign := <-endCh:
			r.logger.Info("Done. Received signature data from participants: %s", sign.Signature)
			return
			// apply transaction
			// r.router.Send(msg, signature)
		case <-r.Quit():
			return
		}
	}
}
