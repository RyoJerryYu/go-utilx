package syncx

import (
	"context"
	"fmt"
	"time"
)

var handlePanicErr func(error) = func(err error) { panic(err) }

func RegisterPanicHandler(f func(error)) {
	handlePanicErr = f
}

type neverDone struct{ context.Context }

func (neverDone) Deadline() (deadline time.Time, ok bool) {
	return
}

func (neverDone) Done() <-chan struct{} {
	return make(chan struct{})
}

func (neverDone) Err() error {
	return nil
}

type GoOption func(ctx context.Context) (context.Context, func())

func Go(ctx context.Context, f func(ctx context.Context), opts ...GoOption) {
	ctx = neverDone{ctx}
	cancels := make([]func(), 0, len(opts))
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				handlePanicErr(err)
			}

			for _, cancel := range cancels {
				cancel()
			}
		}()

		var cancel func()
		for _, opt := range opts {
			ctx, cancel = opt(ctx)
			cancels = append(cancels, cancel)
		}
		f(ctx)
	}()
}

func WithTimeout(d time.Duration) GoOption {
	return func(ctx context.Context) (context.Context, func()) {
		return context.WithTimeout(ctx, d)
	}
}
