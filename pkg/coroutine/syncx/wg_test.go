package syncx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWG(t *testing.T) {
	ctx := context.Background()
	ch := make(chan struct{}, 5)
	wg := WG(ctx)
	for i := 0; i < 10; i++ {
		wg.Go(func(ctx context.Context) {
			ch <- struct{}{}
		})
	}

	assert.Len(t, ch, 5)

	for i := 0; i < 8; i++ {
		_, ok := <-ch
		require.True(t, ok)
	}
	wg.Wait()

	assert.Len(t, ch, 2)

	for i := 0; i < 2; i++ {
		_, ok := <-ch
		require.True(t, ok)
	}
	assert.Len(t, ch, 0)
}
