package slicex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

type RawSliceOperator[T comparable] struct{}

// Implement IOperator for RawSliceOperator

func (RawSliceOperator[T]) Union(a, b []T) []T         { return Union(a, b) }
func (RawSliceOperator[T]) Subtract(a, b []T) []T      { return Subtract(a, b) }
func (RawSliceOperator[T]) Intersect(a, b []T) []T     { return Intersect(a, b) }
func (RawSliceOperator[T]) MergeAll(slices ...[]T) []T { return MergeAll(slices...) }
func (RawSliceOperator[T]) Equal(a []T, b []T) bool    { return Equal(a, b) }
func (RawSliceOperator[T]) Copy(a []T) []T             { return Copy(a) }

var _ icontainer.IOperator[[]string] = RawSliceOperator[string]{}
