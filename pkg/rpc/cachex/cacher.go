package cachex

import (
	"context"
	"time"
)

type Cacher interface {
	Get(ctx context.Context) ([]byte, error)
	Set(ctx context.Context, value []byte, ttl time.Duration) error
}
