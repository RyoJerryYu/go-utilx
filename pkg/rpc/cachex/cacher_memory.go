package cachex

import (
	"context"
	"time"
)

type memoryCacher struct {
	memory []byte
}

func NewMemoryCacher() Cacher {
	return &memoryCacher{
		memory: []byte(""),
	}
}

func (c *memoryCacher) Get(ctx context.Context) ([]byte, error) {
	return c.memory, nil
}

func (c *memoryCacher) Set(ctx context.Context, value []byte, ttl time.Duration) error {
	c.memory = value
	return nil
}
