package syncx

import (
	"context"
	"sync"
)

type wg struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func WG(ctx context.Context) *wg {
	ctx, cancel := context.WithCancel(ctx)
	return &wg{ctx: ctx, cancel: cancel}
}

func (w *wg) Go(f func(ctx context.Context)) {
	w.wg.Add(1)
	Go(w.ctx, func(ctx context.Context) {
		defer w.wg.Done()
		f(ctx)
	})
}

func (w *wg) Wait() {
	w.wg.Wait()
}

func (w *wg) Cancel() {
	w.cancel()
}
