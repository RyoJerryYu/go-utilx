package slicex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

type Operator[T comparable] struct{}

// Implement IOperator for SliceOperator

func (Operator[T]) ForEach(slice Slice[T], fn func(T))   { slice.ForEach(fn) }
func (Operator[T]) Union(a, b Slice[T]) Slice[T]         { return a.Union(b) }
func (Operator[T]) Subtract(a, b Slice[T]) Slice[T]      { return a.Subtract(b) }
func (Operator[T]) Intersect(a, b Slice[T]) Slice[T]     { return a.Intersect(b) }
func (Operator[T]) MergeAll(slices ...Slice[T]) Slice[T] { return SliceMergeAll(slices...) }
func (Operator[T]) Equal(a Slice[T], b Slice[T]) bool    { return a.Equal(b) }
func (Operator[T]) Copy(a Slice[T]) Slice[T]             { return a.Copy() }

var _ icontainer.IOperator[Slice[string]] = Operator[string]{}
