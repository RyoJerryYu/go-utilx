package setx

import "github.com/RyoJerryYu/go-utilx/container/icontainer"

type RawSetOperator[T comparable] struct{}

func (RawSetOperator[T]) Union(a, b map[T]struct{}) map[T]struct{} {
	return RawUnionSet(a, b)
}
func (RawSetOperator[T]) Subtract(a, b map[T]struct{}) map[T]struct{} {
	return RawSubtractSet(a, b)
}
func (RawSetOperator[T]) Intersect(a, b map[T]struct{}) map[T]struct{} {
	return RawIntersectSet(a, b)
}
func (RawSetOperator[T]) Merge(a map[T]struct{}, others ...map[T]struct{}) map[T]struct{} {
	return RawMergeSet(a, others...)
}
func (RawSetOperator[T]) Remove(a map[T]struct{}, others ...map[T]struct{}) map[T]struct{} {
	return RawRemoveSet(a, others...)
}
func (RawSetOperator[T]) Equal(a map[T]struct{}, b map[T]struct{}) bool { return RawEqual(a, b) }
func (RawSetOperator[T]) Copy(a map[T]struct{}) map[T]struct{}          { return RawCopy(a) }

var _ icontainer.IOperator[map[string]struct{}] = RawSetOperator[string]{}
