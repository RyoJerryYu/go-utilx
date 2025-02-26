package cachex

import (
	"context"
	"encoding/json"
	"time"
)

// JSONCacher is a Cacher that stores and retrieves JSON-serializable values.
// It is a wrapper around a Cacher.
type JSONCacher[T any] struct {
	cacher Cacher
}

func NewJSONCacher[T any](cacher Cacher) *JSONCacher[T] {
	return &JSONCacher[T]{
		cacher: cacher,
	}
}

func (c *JSONCacher[T]) Get(ctx context.Context) (*T, error) {
	raw, err := c.cacher.Get(ctx)
	if err != nil {
		return nil, err
	}

	if string(raw) == "" {
		return new(T), nil
	}

	var t T
	err = json.Unmarshal(raw, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (c *JSONCacher[T]) Set(ctx context.Context, value *T, ttl time.Duration) error {
	raw, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.cacher.Set(ctx, raw, ttl)
}
