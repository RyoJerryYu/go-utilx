package setx

import "github.com/RyoJerryYu/go-utilx/container/icontainer"

type RawSetOperator[T comparable] struct{}

// Implement IOperator for RawSetOperator

func (RawSetOperator[T]) Union(a, b map[T]struct{}) map[T]struct{} {
	return Union(a, b)
}
func (RawSetOperator[T]) Subtract(a, b map[T]struct{}) map[T]struct{} {
	return Subtract(a, b)
}
func (RawSetOperator[T]) Intersect(a, b map[T]struct{}) map[T]struct{} {
	return Intersect(a, b)
}
func (RawSetOperator[T]) MergeAll(sets ...map[T]struct{}) map[T]struct{} {
	return MergeAll(sets...)
}
func (RawSetOperator[T]) Equal(a map[T]struct{}, b map[T]struct{}) bool { return Equal(a, b) }
func (RawSetOperator[T]) Copy(a map[T]struct{}) map[T]struct{}          { return Copy(a) }

var _ icontainer.IOperator[map[string]struct{}] = RawSetOperator[string]{}

// Set Specific

func (RawSetOperator[T]) Merge(a map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	return MergeSet(a, other)
}
func (RawSetOperator[T]) Remove(a map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	return RemoveSet(a, other)
}
