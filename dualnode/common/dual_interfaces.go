package common

import (
	"github.com/kardiachain/go-kardia/types"
)

type IWatcher interface {
	Start() error
	Stop() error

	GetLatestDualEvents() ([]types.Log, error)
}
