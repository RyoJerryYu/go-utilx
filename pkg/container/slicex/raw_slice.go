package slicex

// type RawSlice[T any] = []T
// Slice RawSlice is a go type, we cannot add methods to it.

func FromKey[T comparable, V any](in map[T]V) []T {
	out := make([]T, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

func FromValue[T comparable, V any](in map[T]V) []V {
	out := make([]V, 0, len(in))
	for _, v := range in {
		out = append(out, v)
	}
	return out
}

func FromSet[T comparable](in map[T]struct{}) []T {
	out := make([]T, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

func From[T any](in ...T) []T {
	return in
}

func To[T any, I any](in []T, getValue func(T) I) []I {
	out := make([]I, len(in))
	for idx := range in {
		out[idx] = getValue(in[idx])
	}

	return out
}

func MapBy[T any, I comparable](in []T, fn func(T) I) map[I]T {
	out := make(map[I]T, len(in))
	for _, v := range in {
		out[fn(v)] = v
	}
	return out
}

func GroupBy[T any, I comparable](in []T, fn func(T) I) map[I][]T {
	out := make(map[I][]T)
	for _, v := range in {
		k := fn(v)
		out[k] = append(out[k], v)
	}
	return out
}

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

// a & b , in the order of a
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

// a - b , in the order of a
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

// a | b , in the order of a
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

// merge all slices into one. Do not remove duplicates
func MergeAll[T any](slices ...[]T) []T {
	out := make([]T, 0)
	for _, slice := range slices {
		out = append(out, slice...)
	}
	return out
}

// a == b only if a and b have the same elements in the same order
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

func Copy[T any](a []T) []T {
	out := make([]T, len(a))
	copy(out, a)
	return out
}

// a has v
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
