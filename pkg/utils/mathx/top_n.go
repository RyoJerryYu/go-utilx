package mathx

import (
	"github.com/RyoJerryYu/go-utilx/pkg/container/heapx"
)

// SliceTopN returns the N largest elements from the input slice based on the less function.
// The less function should return true if i is considered smaller than j.
// If N is larger than the input slice length, returns all elements.
// The order of elements in the output is not guaranteed.
//
// Example:
//
//	nums := []int{1, 5, 3, 8, 2, 9, 4}
//	less := func(i, j int) bool { return i < j }
//	top3 := SliceTopN(nums, less, 3) // returns [8, 9, 5] (order may vary)
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

// TopN maintains a collection of the N largest elements seen so far.
// Elements can be added one at a time, and the N largest elements
// can be queried at any time.
//
// Example usage:
//
//	topN := NewTopN(3, func(i, j int) bool { return i < j })
//	topN.Push(1)
//	topN.Push(5)
//	topN.Push(3)
//	topN.Push(8)
//	result := topN.Query() // returns [3, 5, 8] (order may vary)
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
