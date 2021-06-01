package common

type IWatcher interface {
	Start() error
	Stop() error
	Watch() error
}

type ChainClient interface {
	//SignTx(tx stypes.TxOutItem, height int64) ([]byte, error)
	//BroadcastTx(_ stypes.TxOutItem, _ []byte) error
	//GetHeight() (int64, error)
	//GetAddress(poolPubKey common.PubKey) string
	//GetAccount(poolPubKey common.PubKey) (common.Account, error)
	//GetChain() common.Chain
	//Start(globalTxsQueue chan stypes.TxIn, globalErrataQueue chan stypes.ErrataBlock)
	//GetConfig() config.ChainConfiguration
	//Stop()
}
