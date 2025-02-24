package timer

import "time"

type ExponentialBackoff struct {
	MaxAttempts    int
	InitialDelay   time.Duration
	MaxDelay       time.Duration
	currentAttempt int
	nextDelay      time.Duration
}

func NewExponentialBackoff(maxElapsedTime time.Duration, maxAttempts int, initialDelay time.Duration, maxDelay time.Duration) *ExponentialBackoff {
	return &ExponentialBackoff{
		MaxAttempts:  maxAttempts,
		InitialDelay: initialDelay,
		MaxDelay:     maxDelay,
	}
}

func (b *ExponentialBackoff) Next() time.Duration {
	if b.currentAttempt == 0 {
		b.currentAttempt++
		b.nextDelay = b.InitialDelay
		return b.nextDelay
	}

	// 计算下一个延迟时间，指数增长
	next := b.nextDelay * 2
	if next > b.MaxDelay {
		next = b.MaxDelay
	}
	b.nextDelay = next
	b.currentAttempt++

	return b.nextDelay
}

func (b *ExponentialBackoff) Reset() {
	b.currentAttempt = 0
	b.nextDelay = 0
}

// Clone returns a new ExponentialBackoff with the same configuration.
func (b *ExponentialBackoff) Clone() Timer {
	return &ExponentialBackoff{
		MaxAttempts:    b.MaxAttempts,
		InitialDelay:   b.InitialDelay,
		MaxDelay:       b.MaxDelay,
		currentAttempt: 0,
		nextDelay:      0,
	}
}
