package tss

import (
	"math/big"

	"github.com/binance-chain/tss-lib/common"
	"github.com/binance-chain/tss-lib/ecdsa/signing"
	"github.com/binance-chain/tss-lib/tss"
	"github.com/gogo/protobuf/proto"

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
	r.signingParties[kcommon.BytesToHash(msg).String()] = party
	r.signingPartyM.Unlock()
	for {
		select {
		case err := <-errCh:
			r.logger.Error("handle signing", "err", err)
			return
		case m := <-outCh:
			r.addOutMsg(m)
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

func (r *Reactor) receiveSigningMsg(msgBytes []byte) {
	msg, err := DecodeMsg(msgBytes)
	if err != nil {
		r.logger.Error("decode signing message err", err)
		return
	}

	_msgBytes, err := proto.Marshal(msg.wire)
	if err != nil {
		r.logger.Error("marshal err", err)
		return
	}

	wireFrom := msg.wire.GetFrom()
	partyId := tss.NewPartyID(wireFrom.Id, wireFrom.Moniker, new(big.Int).SetBytes(wireFrom.Key))
	r.signingParties[msg.ID].UpdateFromBytes(_msgBytes, partyId, true)
}
