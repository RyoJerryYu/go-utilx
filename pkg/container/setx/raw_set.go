package setx

// type RawSet[T comparable] = map[T]struct{}
// Since RawSet is a go type, we cannot add methods to it.

// FromKey creates a set from the keys of a map.
// The resulting set will contain all unique keys from the input map.
func FromKey[T comparable, V any](in map[T]V) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for k := range in {
		out[k] = struct{}{}
	}
	return out
}

// FromSlice creates a set from a slice.
// Duplicate elements in the input slice will only appear once in the set.
func FromSlice[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, v := range in {
		out[v] = struct{}{}
	}
	return out
}

// From creates a set from the given elements.
// Duplicate elements will only appear once in the set.
func From[T comparable](in ...T) map[T]struct{} {
	return FromSlice(in)
}

// ToSlice converts a set to a slice.
// The order of elements in the resulting slice is not guaranteed.
func ToSlice[T comparable](in map[T]struct{}) []T {
	out := make([]T, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

func ForEach[T comparable](in map[T]struct{}, fn func(T)) {
	for k := range in {
		fn(k)
	}
}

// IntersectSlice returns a new set containing elements that exist in both the set and the slice.
// Example:
//
//	set := From(1, 2, 3, 4)
//	slice := []int{3, 4, 5, 6}
//	result := IntersectSlice(set, slice)
//	// result contains: 3, 4
func IntersectSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	out := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := s[v]; ok {
			out[v] = struct{}{}
		}
	}
	return out
}

// IntersectSet returns a new set containing elements that exist in both sets.
// Example:
//
//	set1 := From(1, 2, 3, 4)
//	set2 := From(3, 4, 5, 6)
//	result := IntersectSet(set1, set2)
//	// result contains: 3, 4
func IntersectSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		if _, ok := other[k]; ok {
			out[k] = struct{}{}
		}
	}
	return out
}

// SubtractSlice returns a new set containing elements from the set that do not exist in the slice.
// Example:
//
//	set := From(1, 2, 3, 4)
//	slice := []int{3, 4, 5, 6}
//	result := SubtractSlice(set, slice)
//	// result contains: 1, 2
func SubtractSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	for _, v := range arr {
		delete(out, v)
	}
	return out
}

// SubtractSet returns a new set containing elements from the first set that do not exist in the second set.
// Example:
//
//	set1 := From(1, 2, 3, 4)
//	set2 := From(3, 4, 5, 6)
//	result := SubtractSet(set1, set2)
//	// result contains: 1, 2
func SubtractSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	for k := range other {
		delete(out, k)
	}
	return out
}

// UnionSlice returns a new set containing all unique elements from both the set and the slice.
// Example:
//
//	set := From(1, 2, 3)
//	slice := []int{3, 4, 5}
//	result := UnionSlice(set, slice)
//	// result contains: 1, 2, 3, 4, 5
func UnionSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	for _, v := range arr {
		out[v] = struct{}{}
	}
	return out
}

// UnionSet returns a new set containing all unique elements from both sets.
// Example:
//
//	set1 := From(1, 2, 3)
//	set2 := From(3, 4, 5)
//	result := UnionSet(set1, set2)
//	// result contains: 1, 2, 3, 4, 5
func UnionSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	for k := range other {
		out[k] = struct{}{}
	}
	return out
}

func Intersect[T comparable](a map[T]struct{}, b map[T]struct{}) map[T]struct{} {
	return IntersectSet(a, b)
}
func Subtract[T comparable](a map[T]struct{}, b map[T]struct{}) map[T]struct{} {
	return SubtractSet(a, b)
}
func Union[T comparable](a map[T]struct{}, b map[T]struct{}) map[T]struct{} {
	return UnionSet(a, b)
}

// MergeAll returns a new set containing all unique elements from all input sets.
// Example:
//
//	set1 := From(1, 2)
//	set2 := From(2, 3)
//	set3 := From(3, 4)
//	result := MergeAll(set1, set2, set3)
//	// result contains: 1, 2, 3, 4
func MergeAll[T comparable](sets ...map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for _, s := range sets {
		for k := range s {
			out[k] = struct{}{}
		}
	}
	return out
}

// MergeSlice adds all elements from the slice into the set.
// Note: This function modifies the original set.
// Example:
//
//	set := From(1, 2)
//	slice := []int{2, 3, 4}
//	MergeSlice(set, slice)
//	// set now contains: 1, 2, 3, 4
func MergeSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	for _, v := range arr {
		s[v] = struct{}{}
	}
	return s
}

// MergeSet adds all elements from the other set into the set.
// Note: This function modifies the original set.
// Example:
//
//	set1 := From(1, 2)
//	set2 := From(2, 3, 4)
//	MergeSet(set1, set2)
//	// set1 now contains: 1, 2, 3, 4
func MergeSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	for k := range other {
		s[k] = struct{}{}
	}
	return s
}

// RemoveSlice removes all elements in the slice from the set.
// Note: This function modifies the original set.
// Example:
//
//	set := From(1, 2, 3, 4)
//	slice := []int{2, 4}
//	RemoveSlice(set, slice)
//	// set now contains: 1, 3
func RemoveSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	for _, v := range arr {
		delete(s, v)
	}
	return s
}

// RemoveSet removes all elements in the other set from the set.
// Note: This function modifies the original set.
// Example:
//
//	set1 := From(1, 2, 3, 4)
//	set2 := From(2, 4)
//	RemoveSet(set1, set2)
//	// set1 now contains: 1, 3
func RemoveSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	for k := range other {
		delete(s, k)
	}
	return s
}

// Equal returns true if both sets contain exactly the same elements.
// Example:
//
//	set1 := From(1, 2, 3)
//	set2 := From(1, 2, 3)
//	set3 := From(1, 2, 4)
//	Equal(set1, set2) // returns true
//	Equal(set1, set3) // returns false
func Equal[T comparable](a map[T]struct{}, b map[T]struct{}) bool {
	if len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}

// Copy creates a new set with the same elements as the input set.
func Copy[T comparable](s map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	return out
}

func Has[T comparable](s map[T]struct{}, v T) bool {
	_, ok := s[v]
	return ok
}

func Len[T comparable](s map[T]struct{}) int      { return len(s) }
func IsEmpty[T comparable](s map[T]struct{}) bool { return len(s) == 0 }
func Clear[T comparable](s map[T]struct{}) {
	for k := range s {
		delete(s, k)
	}
}
func Add[T comparable](s map[T]struct{}, vs ...T) map[T]struct{} { return MergeSlice(s, vs) }
func Del[T comparable](s map[T]struct{}, vs ...T) map[T]struct{} { return RemoveSlice(s, vs) }
