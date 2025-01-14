package syncx

import (
	"context"
	"fmt"
	"time"
)

var handlePanicErr func(context.Context, error) = func(ctx context.Context, err error) {
	panic(err)
}

func RegisterPanicHandler(f func(context.Context, error)) {
	handlePanicErr = f
}

type GoOption func(ctx context.Context) (context.Context, func())

func Go(ctx context.Context, f func(ctx context.Context), opts ...GoOption) {
	cancels := make([]func(), 0, len(opts))
	var cancel func()
	for _, opt := range opts {
		ctx, cancel = opt(ctx)
		cancels = append(cancels, cancel)
	}
	go func(ctx context.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				handlePanicErr(ctx, err)
			}

			for _, cancel := range cancels {
				cancel()
			}
		}()

		f(ctx)
	}(ctx)
}

func WithTimeout(d time.Duration) GoOption {
	return func(ctx context.Context) (context.Context, func()) {
		return context.WithTimeout(ctx, d)
	}
}

func WithNoCancel() GoOption {
	return func(ctx context.Context) (context.Context, func()) {
		return context.WithoutCancel(ctx), func() {}
	}
}
