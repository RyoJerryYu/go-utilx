package setx

import (
	"github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"
)

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T]                         { return make(Set[T]) }
func SetFromKey[T comparable, V any](in map[T]V) Set[T] { return FromKey(in) }
func SetFromSlice[T comparable](in []T) Set[T]          { return FromSlice(in) }
func SetFrom[T comparable](in ...T) Set[T]              { return From(in...) }
func (s Set[T]) ToSlice() []T                           { return ToSlice(s) }

func Wrap[T comparable](in map[T]struct{}) Set[T] { return in }
func (s Set[T]) Unwrap() map[T]struct{}           { return s }

func (s Set[T]) IntersectSlice(arr []T) Set[T]    { return IntersectSlice(s, arr) }
func (s Set[T]) IntersectSet(other Set[T]) Set[T] { return IntersectSet(s, other) }
func (s Set[T]) SubtractSlice(arr []T) Set[T]     { return SubtractSlice(s, arr) }
func (s Set[T]) SubtractSet(other Set[T]) Set[T]  { return SubtractSet(s, other) }
func (s Set[T]) UnionSlice(arr []T) Set[T]        { return UnionSlice(s, arr) }
func (s Set[T]) UnionSet(other Set[T]) Set[T]     { return UnionSet(s, other) }

func (s Set[T]) MergeSlice(arr []T) Set[T]     { return MergeSlice(s, arr) }  // Merge other into s, will modify s, return s
func (s Set[T]) RemoveSlice(arr []T) Set[T]    { return RemoveSlice(s, arr) } // Remove  arr from s, will modify s , return s
func (s Set[T]) MergeSet(other Set[T]) Set[T]  { return MergeSet(s, other) }  // Merge other into s, will modify s, return s
func (s Set[T]) RemoveSet(other Set[T]) Set[T] { return RemoveSet(s, other) } // Remove other from s, will modify s , return s

//////
// Operator
//////

func (s Set[T]) Union(other Set[T]) Set[T]     { return s.Union(other) }
func (s Set[T]) Subtract(other Set[T]) Set[T]  { return s.Subtract(other) }
func (s Set[T]) Intersect(other Set[T]) Set[T] { return s.Intersect(other) }
func (s Set[T]) Equal(other Set[T]) bool       { return Equal(s, other) }
func (s Set[T]) Copy() Set[T]                  { return Copy(s) }
func SetMergeAll[T comparable](sets ...Set[T]) Set[T] {
	rawSets := make([]map[T]struct{}, len(sets))
	for i, set := range sets {
		rawSets[i] = set
	}
	return MergeAll(rawSets...)
}

//////
// Set Specific
//////

func (s Set[T]) Merge(other Set[T]) Set[T]  { return s.MergeSet(other) }
func (s Set[T]) Remove(other Set[T]) Set[T] { return s.RemoveSet(other) }

//////
// Container
//////

func (s Set[T]) Len() int      { return Len(s) }
func (s Set[T]) IsEmpty() bool { return IsEmpty(s) }
func (s Set[T]) Clear()        { Clear(s) }
func (s Set[T]) Has(v T) bool  { return Has(s, v) }
func (s Set[T]) Add(vs ...T)   { Add(s, vs...) }
func (s Set[T]) Del(vs ...T)   { Del(s, vs...) }

var _ icontainer.Container[any] = (Set[any])(nil)
