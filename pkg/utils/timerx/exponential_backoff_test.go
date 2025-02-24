package timerx

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestExponentialBackoff(t *testing.T) {
	t.Run("initial delay", func(t *testing.T) {
		eb := NewExponentialBackoff(5, 100*time.Millisecond, 1*time.Second)
		delay := eb.Next()
		require.Equal(t, delay, 100*time.Millisecond)
	})

	t.Run("exponential growth", func(t *testing.T) {
		eb := NewExponentialBackoff(5, 100*time.Millisecond, 1*time.Second)
		expected := []time.Duration{
			100 * time.Millisecond,
			200 * time.Millisecond,
			400 * time.Millisecond,
			800 * time.Millisecond,
			1 * time.Second, // capped at max delay
		}

		for _, exp := range expected {
			got := eb.Next()
			require.Equal(t, got, exp)
		}
	})

	t.Run("reset", func(t *testing.T) {
		eb := NewExponentialBackoff(5, 100*time.Millisecond, 1*time.Second)
		eb.Next() // 100ms
		eb.Next() // 200ms
		eb.Reset()

		delay := eb.Next()
		require.Equal(t, delay, 100*time.Millisecond)
	})

	t.Run("clone", func(t *testing.T) {
		eb := NewExponentialBackoff(5, 100*time.Millisecond, 1*time.Second)
		cloneIface := eb.Clone()
		clone, ok := cloneIface.(*ExponentialBackoff)
		require.True(t, ok)

		require.Equal(t, clone.MaxAttempts, eb.MaxAttempts)
		require.Equal(t, clone.InitialDelay, eb.InitialDelay)
		require.Equal(t, clone.MaxDelay, eb.MaxDelay)

		require.Equal(t, clone.currentAttempt, 0)
		require.Equal(t, clone.nextDelay, time.Duration(0))
	})
}
