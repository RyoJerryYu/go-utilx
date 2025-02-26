package cachex

import (
	"context"
	"time"
)

// RWCacher is a multi-instance friendly cacher with lock.
//
// It is used to cache data that can be modified by multiple instances.
// When modifying the data, it works atomically by obtaining a lock.
// When reading the data, it does not obtain a lock.
// It is a wrapper around a JSONCacher and a LockObtainer.
type RWCacher[T any] struct {
	cache  *JSONCacher[T]
	locker LockObtainer
	ttl    time.Duration
}

func NewRWCacher[T any](cache *JSONCacher[T], locker LockObtainer, ttl time.Duration) *RWCacher[T] {
	return &RWCacher[T]{
		cache:  cache,
		locker: locker,
		ttl:    ttl,
	}
}

func (c *RWCacher[T]) Get(ctx context.Context) (*T, error) {
	return c.cache.Get(ctx)
}

func (c *RWCacher[T]) Modify(ctx context.Context, modify func(value *T) (*T, error)) error {
	lock, err := c.locker.Obtain(ctx)
	if err != nil {
		return err
	}
	defer lock.Release(ctx)

	value, err := c.cache.Get(ctx)
	if err != nil {
		return err
	}

	newValue, err := modify(value)
	if err != nil {
		return err
	}

	return c.cache.Set(ctx, newValue, c.ttl)
}
