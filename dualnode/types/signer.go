package types

type Signer interface {
	Sign(msg []byte) ([]byte, error)
	GetPubKey() []byte
}
