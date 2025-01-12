package syncx

import (
	"context"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	t.Run("test without options", func(t *testing.T) {
		t.Parallel()
		done := make(chan struct{})
		Go(context.Background(), func(ctx context.Context) {
			close(done)
		})
		select {
		case <-done:
		case <-time.After(time.Second):
			t.Fatal("function did not complete in time")
		}
	})

	t.Run("test with timeout", func(t *testing.T) {
		t.Parallel()
		done := make(chan struct{})
		Go(context.Background(), func(ctx context.Context) {
			select {
			case <-ctx.Done():
				close(done)
			case <-time.After(2 * time.Second):
				t.Fatal("context was not cancelled in time")
			}
		}, WithTimeout(10*time.Microsecond))
		select {
		case <-done:
		case <-time.After(3 * time.Second):
			t.Fatal("function did not complete in time")
		}
	})

	t.Run("test panic handling", func(t *testing.T) {
		t.Parallel()
		var panicErr error
		RegisterPanicHandler(func(err error) {
			panicErr = err
		})
		Go(context.Background(), func(ctx context.Context) {
			panic("test panic")
		})
		time.Sleep(1000 * time.Microsecond) // give some time for the goroutine to panic
		if panicErr == nil || panicErr.Error() != "test panic" {
			t.Fatalf("expected panic error, got %v", panicErr)
		}
	})

	t.Run("test multiple options", func(t *testing.T) {
		t.Parallel()
		done := make(chan struct{})
		Go(context.Background(), func(ctx context.Context) {
			select {
			case <-ctx.Done():
				close(done)
			case <-time.After(2 * time.Second):
				t.Fatal("context was not cancelled in time")
			}
		}, WithTimeout(time.Microsecond), WithTimeout(2*time.Microsecond))
		select {
		case <-done:
		case <-time.After(3 * time.Second):
			t.Fatal("function did not complete in time")
		}
	})
}
