package setx

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

type RawOperator[T comparable] struct{}

// Implement IOperator for RawSetOperator

func (RawOperator[T]) ForEach(set map[T]struct{}, fn func(T))         { ForEach(set, fn) }
func (RawOperator[T]) Union(a, b map[T]struct{}) map[T]struct{}       { return Union(a, b) }
func (RawOperator[T]) Subtract(a, b map[T]struct{}) map[T]struct{}    { return Subtract(a, b) }
func (RawOperator[T]) Intersect(a, b map[T]struct{}) map[T]struct{}   { return Intersect(a, b) }
func (RawOperator[T]) MergeAll(sets ...map[T]struct{}) map[T]struct{} { return MergeAll(sets...) }
func (RawOperator[T]) Equal(a map[T]struct{}, b map[T]struct{}) bool  { return Equal(a, b) }
func (RawOperator[T]) Copy(a map[T]struct{}) map[T]struct{}           { return Copy(a) }

var _ icontainer.IOperator[map[string]struct{}] = RawOperator[string]{}

// Set Specific

func (RawOperator[T]) Merge(a map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	return MergeSet(a, other)
}
func (RawOperator[T]) Remove(a map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	return RemoveSet(a, other)
}
