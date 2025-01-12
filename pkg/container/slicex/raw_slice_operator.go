package slicex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

type RawOperator[T comparable] struct{}

// Implement IOperator for RawSliceOperator

func (RawOperator[T]) ForEach(slice []T, fn func(T)) { ForEach(slice, fn) }
func (RawOperator[T]) Union(a, b []T) []T            { return Union(a, b) }
func (RawOperator[T]) Subtract(a, b []T) []T         { return Subtract(a, b) }
func (RawOperator[T]) Intersect(a, b []T) []T        { return Intersect(a, b) }
func (RawOperator[T]) MergeAll(slices ...[]T) []T    { return MergeAll(slices...) }
func (RawOperator[T]) Equal(a []T, b []T) bool       { return Equal(a, b) }
func (RawOperator[T]) Copy(a []T) []T                { return Copy(a) }

var _ icontainer.IOperator[[]string] = RawOperator[string]{}
