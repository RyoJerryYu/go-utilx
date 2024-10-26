package heapx

import (
	"container/heap"

	"golang.org/x/exp/constraints"
)

type LesserThanable[T any] interface {
	LesserThan(T) bool
}
type GreaterThanable[T any] interface {
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

// Heap is a generic heap implement based on container.heap
type Heap[T any] struct {
	in container[T]
}

func (h *Heap[T]) Push(x T) {
	heap.Push(&h.in, x)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(&h.in).(T)
}

func (h *Heap[T]) Remove(i int) T {
	return heap.Remove(&h.in, i).(T)
}

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

func (h *Heap[T]) Snapshot() []T {
	res := make([]T, len(h.in.Data))
	copy(res, h.in.Data)
	return res
}

// New min heap
func NewMin[T constraints.Ordered](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a < b },
	}}
	heap.Init(&h.in)
	return h
}

// New min heap with LessThanable
func NewMinI[T LesserThanable[T]](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a.LesserThan(b) },
	}}
	heap.Init(&h.in)
	return h
}

// New max heap
func NewMax[T constraints.Ordered](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a > b },
	}}
	heap.Init(&h.in)
	return h
}

// New max heap with GreaterThanable
func NewMaxI[T GreaterThanable[T]](data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: func(a, b T) bool { return a.GreaterThan(b) },
	}}
	heap.Init(&h.in)
	return h
}

// New heap with custom less function
// heap top will be the "lessest" element
func NewWith[T any](less func(T, T) bool, data ...T) *Heap[T] {
	h := &Heap[T]{in: container[T]{
		Data: data,
		less: less,
	}}
	heap.Init(&h.in)
	return h
}
