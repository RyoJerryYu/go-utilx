package syncx

import (
	"context"
	"fmt"
	"time"
)

var handlePanicErr func(context.Context, error) = func(ctx context.Context, err error) {
	panic(err)
}

// RegisterPanicHandler sets a custom handler function for panics that occur in
// goroutines started by Go(). The handler receives the context and the error
// that caused the panic.
//
// If no handler is registered, panics will be propagated normally.
func RegisterPanicHandler(f func(context.Context, error)) {
	handlePanicErr = f
}

// GoOption represents a configuration option for the Go function.
// It receives a context and returns a modified context along with
// a cleanup function.
type GoOption func(ctx context.Context) (context.Context, func())

// Go launches a goroutine with panic recovery and optional context modifications.
// The provided function f is executed in a new goroutine with the given context.
// Any panics in the goroutine will be caught and handled by the registered panic handler.
//
// Example usage:
//
//	Go(ctx, func(ctx context.Context) {
//	    // Your goroutine logic here
//	}, WithTimeout(5*time.Second))
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

// WithTimeout returns a GoOption that adds a timeout to the context.
// The goroutine will be cancelled after the specified duration.
func WithTimeout(d time.Duration) GoOption {
	return func(ctx context.Context) (context.Context, func()) {
		return context.WithTimeout(ctx, d)
	}
}

// WithNoCancel returns a GoOption that creates a context that cannot be cancelled.
// This is useful when you want the goroutine to complete regardless of the
// parent context's cancellation.
func WithNoCancel() GoOption {
	return func(ctx context.Context) (context.Context, func()) {
		return context.WithoutCancel(ctx), func() {}
	}
}
