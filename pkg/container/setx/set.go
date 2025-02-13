package setx

import (
	"github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"
)

// Set is a generic set type implemented using a map with empty struct values.
// Type parameter T must be comparable to support set operations.
type Set[T comparable] map[T]struct{}

// New creates a new empty Set.
func New[T comparable]() Set[T]                         { return make(Set[T]) }
func SetFromKey[T comparable, V any](in map[T]V) Set[T] { return FromKey(in) }
func SetFromSlice[T comparable](in []T) Set[T]          { return FromSlice(in) }
func SetFrom[T comparable](in ...T) Set[T]              { return From(in...) }
func (s Set[T]) ToSlice() []T                           { return ToSlice(s) }

// Wrap converts a raw map[T]struct{} to a Set[T].
// This is a zero-cost conversion as it just changes the type.
func Wrap[T comparable](in map[T]struct{}) Set[T] { return in }

// Unwrap converts a Set[T] back to a raw map[T]struct{}.
// This is a zero-cost conversion as it just changes the type.
func (s Set[T]) Unwrap() map[T]struct{} { return s }

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

func (s Set[T]) Union(other Set[T]) Set[T]     { return Union(s, other) }
func (s Set[T]) Subtract(other Set[T]) Set[T]  { return Subtract(s, other) }
func (s Set[T]) Intersect(other Set[T]) Set[T] { return Intersect(s, other) }
func (s Set[T]) Equal(other Set[T]) bool       { return Equal(s, other) }
func (s Set[T]) Copy() Set[T]                  { return Copy(s) }

// SetMergeAll combines multiple Sets into a single Set.
// Duplicate elements will only appear once in the result.
// Example:
//
//	set1 := SetFrom(1, 2)
//	set2 := SetFrom(2, 3)
//	set3 := SetFrom(3, 4)
//	result := SetMergeAll(set1, set2, set3)
//	// result contains: 1, 2, 3, 4
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

func (s Set[T]) Len() int           { return Len(s) }
func (s Set[T]) IsEmpty() bool      { return IsEmpty(s) }
func (s Set[T]) Clear()             { Clear(s) }
func (s Set[T]) ForEach(fn func(T)) { ForEach(s, fn) }
func (s Set[T]) Has(v T) bool       { return Has(s, v) }
func (s Set[T]) Add(vs ...T)        { Add(s, vs...) }
func (s Set[T]) Del(vs ...T)        { Del(s, vs...) }

var _ icontainer.Container[any] = (Set[any])(nil)
