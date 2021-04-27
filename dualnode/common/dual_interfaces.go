package common

type IWatcher interface {
	Start() error
	Stop() error
	Watch() error
}
