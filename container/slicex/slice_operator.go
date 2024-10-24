package slicex

import "github.com/RyoJerryYu/go-utilx/container/icontainer"

type SliceOperator[T comparable] struct{}

// Implement IOperator for SliceOperator

func (SliceOperator[T]) Union(a, b Slice[T]) Slice[T]         { return a.Union(b) }
func (SliceOperator[T]) Subtract(a, b Slice[T]) Slice[T]      { return a.Subtract(b) }
func (SliceOperator[T]) Intersect(a, b Slice[T]) Slice[T]     { return a.Intersect(b) }
func (SliceOperator[T]) MergeAll(slices ...Slice[T]) Slice[T] { return SliceMergeAll(slices...) }
func (SliceOperator[T]) Equal(a Slice[T], b Slice[T]) bool    { return a.Equal(b) }
func (SliceOperator[T]) Copy(a Slice[T]) Slice[T]             { return a.Copy() }

var _ icontainer.IOperator[Slice[string]] = SliceOperator[string]{}
