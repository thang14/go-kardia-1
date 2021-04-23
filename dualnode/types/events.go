package types

import (
	"time"

	"github.com/kardiachain/go-kardia/types"
)

type DualEvent struct {
	Source      int64       `json:"source"`
	Destination int64       `json:"destination"`
	Arguments   interface{} `json:"arguments"`
	Timestamp   time.Time   `json:"timestamp"`

	RawName string    `json:"rawName"`
	Raw     types.Log `json:"raw"`
}
