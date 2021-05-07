package types

type TransferOwnershipMsg struct {
	Owner []byte
}

type AddTokenMsg struct {
	Symbol    [32]byte
	TokenAddr []byte
	Locktype  uint8
}

type RemoveToken struct {
	Symbol [32]byte
}
