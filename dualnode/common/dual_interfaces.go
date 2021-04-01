package common

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	dualCfg "github.com/kardiachain/go-kardia/dualnode/config"
	dualTypes "github.com/kardiachain/go-kardia/dualnode/types"
	"github.com/kardiachain/go-kardia/lib/abi"
	"github.com/kardiachain/go-kardia/lib/common"
)

type IWatcher interface {
	Start() error
	Stop() error

	GetLatestDualEvents() ([]*dualTypes.DualEvent, error)
	GetDualEventsChannel() chan *dualTypes.DualEvent
}

// UnpackDualEventIntoMap unpacks indexed and unindexed arguments of a retrieved log into a map[string]interface{}.
func UnpackDualEventIntoMap(a *abi.ABI, dualEvent *dualTypes.DualEvent, src int64) (interface{}, int64, int64, error) {
	out := make(map[string]interface{})
	data, err := hex.DecodeString(strings.TrimPrefix(dualEvent.Data.String(), "0x"))
	if err != nil {
		return nil, -1, -1, err
	}
	// unpacking unindexed arguments
	if len(data) > 0 {
		if err := a.UnpackIntoMap(out, dualEvent.RawName, data); err != nil {
			return nil, -1, -1, err
		}
	}
	// unpacking indexed arguments
	var indexed abi.Arguments
	for _, arg := range a.Events[dualEvent.RawName].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	topics := make([]common.Hash, len(dualEvent.Topics)-1)
	for i, topic := range dualEvent.Topics[1:] { // exclude the eventID (dualEvent.Topic[0])
		topics[i] = topic
	}
	err = abi.ParseTopicsIntoMap(out, indexed, topics)
	if err != nil {
		return nil, -1, -1, fmt.Errorf("cannot parse indexed arguments to map: %v", err)
	}

	// fulfill all arguments of a dual event by its type
	raw, err := json.Marshal(out)
	if err != nil {
		return nil, -1, -1, fmt.Errorf("cannot marshal event data: %v", err)
	}
	switch dualEvent.RawName {
	case dualCfg.LockEventRawName:
		var lockEvent *dualTypes.LockParams
		err = json.Unmarshal(raw, &lockEvent)
		if err != nil {
			return nil, -1, -1, fmt.Errorf("cannot unmarshal to a lock event: %v lockEvent: %+v", err, lockEvent)
		}
		return lockEvent, src, lockEvent.Destination.Int64(), nil
	case dualCfg.UnlockEventRawName:
		var unlockEvent *dualTypes.UnlockParams
		err = json.Unmarshal(raw, &unlockEvent)
		if err != nil {
			return nil, -1, -1, fmt.Errorf("cannot unmarshal to a unlock event: %v unlockEvent: %+v", err, unlockEvent)
		}
		return unlockEvent, unlockEvent.Source.Int64(), unlockEvent.Destination.Int64(), nil
	case dualCfg.AddValidatorEventRawName:
	case dualCfg.RemoveValidatorEventRawName:
	}
	return nil, -1, -1, fmt.Errorf("dual event doesn't belong to any type")
}
