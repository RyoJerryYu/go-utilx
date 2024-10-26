package mathx

import (
	"github.com/RyoJerryYu/go-utilx/pkg/container/heapx"
)

// Order is not guaranteed
// will return the Largest N based on less
func SliceTopN[T any](in []T, less func(T, T) bool, N int) []T {
	if N < 0 {
		return nil
	}

	out := make([]T, 0, N)
	if len(in) < N {
		out = append(out, in...)
		return out
	}

	h := heapx.NewWith(less, in[:N]...)
	for i := N; i < len(in); i++ {
		h.Push(in[i])
		h.Pop()
	}

	return h.Data()
}

type TopN[T any] struct {
	n int
	h *heapx.Heap[T]
}

// thread unsafe!
// Query will return the Largest N based on less
func NewTopN[T any](N int, less func(T, T) bool) *TopN[T] {
	return &TopN[T]{
		n: N,
		h: heapx.NewWith(less),
	}
}

func (t *TopN[T]) Push(item T) {
	t.h.Push(item)

	if t.h.Len() > t.n {
		t.h.Pop()
	}
}

// Order is not guaranteed
func (t *TopN[T]) Query() []T {
	return t.h.Snapshot()
}
