package dualnode

type API struct {
	service *Service
}

func (api *API) GetDepositSignatures(hash []byte) ([][]byte, error) {
	return nil, nil
}

// NewPublicTransactionAPI is a constructor of PublicTransactionAPI
func NewAPI(service *Service) *API {
	return &API{service}
}
