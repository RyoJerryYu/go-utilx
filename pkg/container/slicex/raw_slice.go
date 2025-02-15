package slicex

// type RawSlice[T any] = []T
// Slice RawSlice is a go type, we cannot add methods to it.

// FromKey extracts all keys from a map into a slice.
// The order of elements in the resulting slice is not guaranteed.
func FromKey[T comparable, V any](in map[T]V) []T {
	out := make([]T, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// FromValue extracts all values from a map into a slice.
// The order of elements in the resulting slice is not guaranteed.
func FromValue[T comparable, V any](in map[T]V) []V {
	out := make([]V, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

// FromSet converts a set (implemented as map[T]struct{}) into a slice.
// The order of elements in the resulting slice is not guaranteed.
func FromSet[T comparable](in map[T]struct{}) []T {
	out := make([]T, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// From creates a slice from given elements, preserving their order.
func From[T any](in ...T) []T {
	return in
}

// To transforms a slice of type T into a slice of type I using the provided mapping function.
// The order of elements is preserved.
// Example:
//
//	input := []int{1, 2, 3}
//	output := To(input, func(i int) string { return strconv.Itoa(i) })
//	// output = []string{"1", "2", "3"}
func To[T any, I any](in []T, getValue func(T) I) []I {
	out := make([]I, len(in))
	for idx := range in {
		out[idx] = getValue(in[idx])
	}

	return out
}

// MapBy creates a map from a slice using the provided function to generate keys.
// If multiple elements generate the same key, the last element will be kept.
// Example:
//
//	type User struct { ID int; Name string }
//	users := []User{{1, "Alice"}, {2, "Bob"}}
//	userMap := MapBy(users, func(u User) int { return u.ID })
//	// userMap = map[int]User{1: {1, "Alice"}, 2: {2, "Bob"}}
func MapBy[T any, I comparable](in []T, fn func(T) I) map[I]T {
	out := make(map[I]T, len(in))
	for _, v := range in {
		out[fn(v)] = v
	}
	return out
}

// GroupBy creates a map of slices from a slice using the provided function to generate keys.
// Elements that generate the same key will be grouped together in the same slice.
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	grouped := GroupBy(numbers, func(n int) string { return fmt.Sprintf("%d", n%2) })
//	// grouped = map[string][]int{"0": {2, 4}, "1": {1, 3}}
func GroupBy[T any, I comparable](in []T, fn func(T) I) map[I][]T {
	out := make(map[I][]T)
	for _, v := range in {
		k := fn(v)
		out[k] = append(out[k], v)
	}
	return out
}

// ToSet converts a slice into a set (implemented as map[T]struct{}).
// Duplicates in the input slice will be removed.
func ToSet[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, v := range in {
		out[v] = struct{}{}
	}
	return out
}

func ForEach[T any](in []T, fn func(T)) {
	for _, v := range in {
		fn(v)
	}
}

func Filter[T any](in []T, fn func(T) bool) []T {
	out := make([]T, 0, len(in))
	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

// Intersect returns a new slice containing elements that exist in both slices.
// The order of elements in the result follows the order in slice a.
// Example:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	result := Intersect(a, b)
//	// result = []int{3, 4}
func Intersect[T comparable](a []T, b []T) []T {
	set := ToSet(b)
	out := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := set[v]; ok {
			out = append(out, v)
		}
	}
	return out
}

// Subtract returns a new slice containing elements from slice a that do not exist in slice b.
// The order of elements in the result follows the order in slice a.
// Example:
//
//	a := []int{1, 2, 3, 4}
//	b := []int{3, 4, 5, 6}
//	result := Subtract(a, b)
//	// result = []int{1, 2}
func Subtract[T comparable](a []T, b []T) []T {
	set := ToSet(b)
	out := make([]T, 0, len(a))
	for _, v := range a {
		if _, ok := set[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

// Union returns a new slice containing all unique elements from both slices.
// Elements from slice a come first, followed by elements from slice b that don't exist in slice a.
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{3, 4, 5}
//	result := Union(a, b)
//	// result = []int{1, 2, 3, 4, 5}
func Union[T comparable](a []T, b []T) []T {
	set := ToSet(a)
	out := a[:]
	for _, v := range b {
		if _, ok := set[v]; !ok {
			out = append(out, v)
		}
	}
	return out
}

// MergeAll concatenates all given slices into a single slice.
// The order of elements is preserved and duplicates are not removed.
// Example:
//
//	a := []int{1, 2}
//	b := []int{2, 3}
//	c := []int{3, 4}
//	result := MergeAll(a, b, c)
//	// result = []int{1, 2, 2, 3, 3, 4}
func MergeAll[T any](slices ...[]T) []T {
	out := make([]T, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}

// Equal returns true if both slices have the same length and contain the same elements in the same order.
// Example:
//
//	a := []int{1, 2, 3}
//	b := []int{1, 2, 3}
//	c := []int{3, 2, 1}
//	Equal(a, b) // returns true
//	Equal(a, c) // returns false
func Equal[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Copy creates a new slice with the same elements as the input slice.
// The returned slice has its own underlying array.
func Copy[T any](a []T) []T {
	out := make([]T, len(a))
	copy(out, a)
	return out
}

// Has returns true if the slice contains the given element.
func Has[T comparable](a []T, v T) bool {
	for _, e := range a {
		if e == v {
			return true
		}
	}
	return false
}

func Len[T any](a []T) int      { return len(a) }
func IsEmpty[T any](a []T) bool { return len(a) == 0 }
