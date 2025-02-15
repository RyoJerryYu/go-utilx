package slicex

import "github.com/RyoJerryYu/go-utilx/pkg/container/icontainer"

// Slice is a generic slice type that implements various container operations.
// Type parameter T must be comparable to support operations like Union and Intersect.
type Slice[T comparable] []T

// New creates a new empty Slice.
func New[T comparable]() Slice[T]                                    { return make(Slice[T], 0) }
func SliceFromKey[T comparable, V any](in map[T]V) Slice[T]          { return FromKey(in) }
func SliceFromValue[T comparable, V comparable](in map[T]V) Slice[V] { return FromValue(in) }
func SliceFromSet[T comparable](in map[T]struct{}) Slice[T]          { return FromSet(in) }
func SliceFrom[T comparable](in ...T) Slice[T]                       { return From(in...) }
func (s Slice[T]) ToSet() map[T]struct{}                             { return ToSet(s) }
func (s Slice[T]) Filter(fn func(T) bool) Slice[T]                   { return Filter(s, fn) }

//////
// Operator
//////

func (s Slice[T]) Intersect(other Slice[T]) Slice[T] { return Intersect(s, other) }
func (s Slice[T]) Subtract(other Slice[T]) Slice[T]  { return Subtract(s, other) }
func (s Slice[T]) Union(other Slice[T]) Slice[T]     { return Union(s, other) }
func (s Slice[T]) Equal(other Slice[T]) bool         { return Equal(s, other) }
func (s Slice[T]) Copy() Slice[T]                    { return Copy(s) }

// SliceMergeAll combines multiple Slices into a single Slice.
// The order of elements is preserved and duplicates are not removed.
// Example:
//
//	a := SliceFrom(1, 2)
//	b := SliceFrom(2, 3)
//	result := SliceMergeAll(a, b)
//	// result = Slice{1, 2, 2, 3}
func SliceMergeAll[T comparable](others ...Slice[T]) Slice[T] {
	rawSlices := make([][]T, len(others))
	for i, s := range others {
		rawSlices[i] = s
	}
	return MergeAll(rawSlices...)
}

//////
// Slice Specific
//////

func (s Slice[T]) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

//////
// Container
//////

func (s Slice[T]) Len() int           { return Len(s) }
func (s Slice[T]) IsEmpty() bool      { return IsEmpty(s) }
func (s *Slice[T]) Clear()            { *s = make(Slice[T], 0) }
func (s Slice[T]) ForEach(fn func(T)) { ForEach(s, fn) }
func (s Slice[T]) Has(v T) bool       { return Has(s, v) }
func (s *Slice[T]) Add(vs ...T)       { *s = append(*s, vs...) }
func (s *Slice[T]) Del(vs ...T)       { *s = s.Subtract(vs) }

var _ icontainer.Container[string] = (*Slice[string])(nil)
