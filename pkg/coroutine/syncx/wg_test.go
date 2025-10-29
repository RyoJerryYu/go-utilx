package syncx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWG(t *testing.T) {
	ctx := context.Background()
	const (
		chSize    = 5
		total     = 10
		batchSize = 8
	)
	ch := make(chan struct{}, chSize)
	wg1 := WG(ctx)
	for i := 0; i < chSize; i++ {
		wg1.Go(func(ctx context.Context) {
			ch <- struct{}{}
		})
	}
	wg1.Wait()

	assert.Len(t, ch, chSize)

	wg2 := WG(ctx)
	for i := 0; i < total-chSize; i++ {
		wg2.Go(func(ctx context.Context) {
			ch <- struct{}{}
		})
	}
	for i := 0; i < batchSize; i++ {
		_, ok := <-ch
		require.True(t, ok)
	}
	wg2.Wait()

	assert.Len(t, ch, total-batchSize)

	for i := 0; i < total-batchSize; i++ {
		_, ok := <-ch
		require.True(t, ok)
	}
	assert.Len(t, ch, 0)
}
