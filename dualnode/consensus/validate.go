package consensus

import (
	"fmt"

	dproto "github.com/kardiachain/go-kardia/proto/kardiachain/dualnode"
)

func validateVote(vote *dproto.Vote) error {
	if vote.Addr == nil {
		return fmt.Errorf("required addr")
	}
	if vote.Hash == nil {
		return fmt.Errorf("required addr")
	}
	if len(vote.Signature) != 65 {
		return fmt.Errorf("invalid signature")
	}
	return nil
}
