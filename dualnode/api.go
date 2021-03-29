package dualnode

type BridgeAPI struct {
	service *Service
}

func (api *BridgeAPI) Signs(chainID int64, depositID int64) ([][]byte, error) {
	return api.service.Signs(chainID, depositID), nil
}

// BriNewBridgeAPIdgeAPI is a constructor of BridgeAPI
func NewBridgeAPI(service *Service) *BridgeAPI {
	return &BridgeAPI{service}
}
