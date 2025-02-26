package cachex

import "context"

type Lock interface {
	Release(ctx context.Context) error
}

type LockObtainer interface {
	Obtain(ctx context.Context) (Lock, error)
}
