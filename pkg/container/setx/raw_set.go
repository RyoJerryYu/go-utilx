package setx

// type RawSet[T comparable] = map[T]struct{}
// Since RawSet is a go type, we cannot add methods to it.

func FromKey[T comparable, V any](in map[T]V) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for k := range in {
		out[k] = struct{}{}
	}
	return out
}

func FromSlice[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, v := range in {
		out[v] = struct{}{}
	}
	return out
}

func From[T comparable](in ...T) map[T]struct{} {
	return FromSlice(in)
}

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

// s & arr
func IntersectSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	out := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := s[v]; ok {
			out[v] = struct{}{}
		}
	}
	return out
}

// s & other
func IntersectSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		if _, ok := other[k]; ok {
			out[k] = struct{}{}
		}
	}
	return out
}

// s - arr
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

// s - other
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

// s | arr
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

// s | other
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

func MergeAll[T comparable](sets ...map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for _, s := range sets {
		for k := range s {
			out[k] = struct{}{}
		}
	}
	return out
}

// Merge arr into s, will modify s, return s
func MergeSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	for _, v := range arr {
		s[v] = struct{}{}
	}
	return s
}

// Merge other into s, will modify s, return s
func MergeSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	for k := range other {
		s[k] = struct{}{}
	}
	return s
}

// Remove arr from s, will modify s, return s
func RemoveSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	for _, v := range arr {
		delete(s, v)
	}
	return s
}

// Remove other from s, will modify s, return s
func RemoveSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	for k := range other {
		delete(s, k)
	}
	return s
}

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
