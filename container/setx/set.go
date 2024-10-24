package setx

import "github.com/RyoJerryYu/go-utilx/container/icontainer"

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T]                      { return NewRaw[T]() }
func FromKey[T comparable, V any](in map[T]V) Set[T] { return RawFromKey(in) }
func FromSlice[T comparable](in []T) Set[T]          { return RawFromSlice(in) }
func From[T comparable](in ...T) Set[T]              { return RawFrom(in...) }
func (s Set[T]) ToSlice() []T                        { return RawToSlice(s) }

func Wrap[T comparable](in map[T]struct{}) Set[T] { return in }
func (s Set[T]) Unwrap() map[T]struct{}           { return s }

func (s Set[T]) IntersectSlice(arr []T) Set[T]    { return RawIntersectSlice(s, arr) }
func (s Set[T]) IntersectSet(other Set[T]) Set[T] { return RawIntersectSet(s, other) }
func (s Set[T]) SubtractSlice(arr []T) Set[T]     { return RawSubtractSlice(s, arr) }
func (s Set[T]) SubtractSet(other Set[T]) Set[T]  { return RawSubtractSet(s, other) }
func (s Set[T]) UnionSlice(arr []T) Set[T]        { return RawUnionSlice(s, arr) }
func (s Set[T]) UnionSet(other Set[T]) Set[T]     { return RawUnionSet(s, other) }

func (s Set[T]) MergeSlice(arrs ...[]T) Set[T]  { return RawMergeSlice(s, arrs...) }  // Merge other into s, will modify s, return s
func (s Set[T]) RemoveSlice(arrs ...[]T) Set[T] { return RawRemoveSlice(s, arrs...) } // Remove  arr from s, will modify s , return s

// Merge other into s, will modify s, return s
func (s Set[T]) MergeSet(others ...Set[T]) Set[T] {
	rawOthers := make([]map[T]struct{}, len(others))
	for i, other := range others {
		rawOthers[i] = other
	}
	return RawMergeSet(s, rawOthers...)
}

// Remove other from s, will modify s , return s
func (s Set[T]) RemoveSet(others ...Set[T]) Set[T] {
	rawOthers := make([]map[T]struct{}, len(others))
	for i, other := range others {
		rawOthers[i] = other
	}
	return RawRemoveSet(s, rawOthers...)
}

//////
// Operator
//////

func (s Set[T]) Union(other Set[T]) Set[T]     { return s.UnionSet(other) }
func (s Set[T]) Subtract(other Set[T]) Set[T]  { return s.SubtractSet(other) }
func (s Set[T]) Intersect(other Set[T]) Set[T] { return s.IntersectSet(other) }
func (s Set[T]) Merge(other Set[T]) Set[T]     { return s.MergeSet(other) }
func (s Set[T]) Remove(other Set[T]) Set[T]    { return s.RemoveSet(other) }
func (s Set[T]) Equal(other Set[T]) bool       { return RawEqual(s, other) }
func (s Set[T]) Copy() Set[T]                  { return RawCopy(s) }

//////
// Container
//////

func (s Set[T]) Len() int      { return RawLen(s) }
func (s Set[T]) IsEmpty() bool { return RawIsEmpty(s) }
func (s Set[T]) Clear()        { RawClear(s) }
func (s Set[T]) Has(v T) bool  { return RawHas(s, v) }
func (s Set[T]) Add(vs ...T)   { RawAdd(s, vs...) }
func (s Set[T]) Del(vs ...T)   { RawDel(s, vs...) }

var _ icontainer.Container[any] = (Set[any])(nil)
