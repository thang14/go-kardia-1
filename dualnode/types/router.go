package types

type Router interface {
	Send(destChainId int64, msgI interface{}) error
}

type Chain interface {
	Receive(msgI interface{}) error
}

type router struct {
	chains map[int64]Chain
}

func NewRouter() Router {
	return &router{}
}

func (r *router) Send(destChainId int64, msgI interface{}) error {
	return r.chains[destChainId].Receive(msgI)
}
