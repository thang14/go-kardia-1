package dualnode

import "github.com/kardiachain/go-kardia/lib/common"

type BridgeAPI struct {
	service *Service
}

func (api *BridgeAPI) Signs(hash common.Hash) ([][]byte, error) {
	return api.service.Signs(hash), nil
}

// BriNewBridgeAPIdgeAPI is a constructor of BridgeAPI
func NewBridgeAPI(service *Service) *BridgeAPI {
	return &BridgeAPI{service}
}
