package kardiachain

import "fmt"

type Watcher struct {
	quit   chan struct{}
	router Router
}

func newWatcher() *Watcher {
	return &Watcher{
		quit: make(chan struct{}, 1),
	}
}

func (w *Watcher) SetRouter(r Router) {
	w.router = r
}

func (w *Watcher) start() error {
	go func() {
		if err := w.watch(); err != nil {
			fmt.Printf("watch blocks error: %s", err)
		}
	}()
	return nil
}

func (w *Watcher) stop() error {
	close(w.quit)
	return nil
}

func (w *Watcher) watch() error {
	for {
		select {
		case <-w.quit:
			return nil
		default:

		}
	}
}
