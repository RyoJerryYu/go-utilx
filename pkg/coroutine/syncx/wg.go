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

// WG creates a new wait group with context support.
// It combines the functionality of sync.WaitGroup with context cancellation.
// The returned wait group will be cancelled when the provided context is cancelled.
//
// Example usage:
//
//	wg := WG(ctx)
//	for i := 0; i < 5; i++ {
//	    wg.Go(func(ctx context.Context) {
//	        // Your concurrent work here
//	    })
//	}
//	wg.Wait() // Wait for all goroutines to complete
func WG(ctx context.Context) *wg {
	ctx, cancel := context.WithCancel(ctx)
	return &wg{ctx: ctx, cancel: cancel}
}

// Go adds a new goroutine to the wait group. The provided function will be
// executed in a new goroutine, and the wait group's counter will be automatically
// managed. The goroutine will receive a context that is cancelled when either
// the wait group's context is cancelled or Cancel() is called.
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

// Cancel cancels all goroutines started by this wait group by cancelling
// the internal context. Note that this does not wait for the goroutines
// to complete - use Wait() for that.
func (w *wg) Cancel() {
	w.cancel()
}
