package tss

import (
	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	"github.com/binance-chain/tss-lib/tss"

	kcommon "github.com/kardiachain/go-kardia/lib/common"
)

func (r *Reactor) Sign(msg []byte) ([]byte, error) {
	peerCtx := tss.NewPeerContext(r.state.partyIDs)
	totalParties := len(r.state.partyIDs)
	params := tss.NewParameters(peerCtx, r.state.localPartyID, totalParties, totalParties/2+1)

	outCh := make(chan tss.Message)
	endCh := make(chan common.SignatureData)
	errCh := make(chan *tss.Error)

	party := signing.NewLocalParty(kcommon.HashToInt(msg), params, r.state.localPartySaveData, outCh, endCh)
	r.addParty(kcommon.BytesToHash(msg), party)
	for {
		select {
		case err := <-errCh:
			r.logger.Error("handle signing", "err", err)
			return nil, err
		case m := <-outCh:
			r.addOutMsg(m)
		case signData := <-endCh:
			r.logger.Info("Done. Received signature data from participants: %s", signData.Signature)
			return signData.Signature, nil
		case <-r.Quit():
			return nil, nil
		}
	}
}
