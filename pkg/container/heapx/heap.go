package heapx

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

// LesserThanable represents types that can be compared with "less than" operation.
// Types implementing this interface can be used in min heaps.
type LesserThanable[T any] interface {
	// LesserThan returns true if the current value is less than the given value.
	LesserThan(T) bool
}

// GreaterThanable represents types that can be compared with "greater than" operation.
// Types implementing this interface can be used in max heaps.
type GreaterThanable[T any] interface {
	// GreaterThan returns true if the current value is greater than the given value.
	GreaterThan(T) bool
}

// container is a heap.Interface implementation
type container[T any] struct {
	Data []T
	less func(T, T) bool
}

func (h container[T]) Len() int           { return len(h.Data) }
func (h container[T]) Less(i, j int) bool { return h.less(h.Data[i], h.Data[j]) }
func (h container[T]) Swap(i, j int)      { h.Data[i], h.Data[j] = h.Data[j], h.Data[i] }

func (h *container[T]) Push(x any) {
	h.Data = append(h.Data, x.(T))
}

func (h *container[T]) Pop() interface{} {
	n := len(h.Data)
	x := h.Data[n-1]
	h.Data = h.Data[0 : n-1]
	return x
}

var _ heap.Interface = (*container[any])(nil)

// Heap is a generic heap implementation based on container/heap.
// It supports both min and max heaps, and can work with custom comparison functions.
// The heap maintains the property that the top element is always the "least" according to the comparison function.
type Heap[T any] struct {
	in container[T]
}

// Push adds an element to the heap.
// The heap property is maintained after the operation.
func (h *Heap[T]) Push(x T) {
	heap.Push(&h.in, x)
}

// Pop removes and returns the top element from the heap.
// For min heaps, this is the smallest element.
// For max heaps, this is the largest element.
// Panics if the heap is empty.
func (h *Heap[T]) Pop() T {
	return heap.Pop(&h.in).(T)
}

// Remove removes and returns the element at index i.
// The heap property is maintained after the operation.
// Panics if i is out of range.
func (h *Heap[T]) Remove(i int) T {
	return heap.Remove(&h.in, i).(T)
}

// Fix re-establishes the heap ordering after the element at index i has been changed.
// This is useful when the element's value has been modified in a way that could violate the heap property.
// Panics if i is out of range.
func (h *Heap[T]) Fix(i int) {
	heap.Fix(&h.in, i)
}

func (h *Heap[T]) Len() int {
	return h.in.Len()
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Len() == 0
}

func (h *Heap[T]) Data() []T {
	return h.in.Data
}

// Snapshot returns a copy of the heap's current elements.
// The returned slice is independent of the heap and can be safely modified.
func (h *Heap[T]) Snapshot() []T {
	res := make([]T, len(h.in.Data))
	copy(res, h.in.Data)
	return res
}

// NewMin creates a new min heap for ordered types.
// Example:
//
//	h := NewMin(5, 3, 1, 4, 2)
//	h.Pop() // returns 1
//	h.Pop() // returns 2
func NewMin[T constraints.Ordered](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a < b },
	}}
	heap.Init(&h.in)
	return h
}

// NewMinI creates a new min heap for types implementing LesserThanable.
// Example:
//
//	type MyInt int
//	func (a MyInt) LesserThan(b MyInt) bool { return a < b }
//	h := NewMinI(MyInt(5), MyInt(3), MyInt(1))
//	h.Pop() // returns MyInt(1)
func NewMinI[T LesserThanable[T]](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a.LesserThan(b) },
	}}
	heap.Init(&h.in)
	return h
}

// NewMax creates a new max heap for ordered types.
// Example:
//
//	h := NewMax(1, 3, 5, 2, 4)
//	h.Pop() // returns 5
//	h.Pop() // returns 4
func NewMax[T constraints.Ordered](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a > b },
	}}
	heap.Init(&h.in)
	return h
}

// NewMaxI creates a new max heap for types implementing GreaterThanable.
// Example:
//
//	type MyInt int
//	func (a MyInt) GreaterThan(b MyInt) bool { return a > b }
//	h := NewMaxI(MyInt(1), MyInt(3), MyInt(5))
//	h.Pop() // returns MyInt(5)
func NewMaxI[T GreaterThanable[T]](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a.GreaterThan(b) },
	}}
	heap.Init(&h.in)
	return h
}

// NewWith creates a new heap with a custom comparison function.
// The less function should return true if a is considered less than b.
// The heap will maintain the property that the top element is the "least" according to this function.
// Example:
//
//	// Create a min heap based on absolute values
//	h := NewWith(func(a, b int) bool {
//	    return abs(a) < abs(b)
//	}, -5, 2, -3, 4, -1)
//	h.Pop() // returns -1 (smallest absolute value)
func NewWith[T any](less func(T, T) bool, data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: less,
	}}
	heap.Init(&h.in)
	return h
}
