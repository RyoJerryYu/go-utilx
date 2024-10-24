package setx

// type RawSet[T comparable] = map[T]struct{}
// Since RawSet is a go type, we cannot add methods to it.

func NewRaw[T comparable]() map[T]struct{} {
	return make(map[T]struct{})
}

func RawFromKey[T comparable, V any](in map[T]V) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for k := range in {
		out[k] = struct{}{}
	}
	return out
}

func RawFromSlice[T comparable](in []T) map[T]struct{} {
	out := make(map[T]struct{}, len(in))
	for _, v := range in {
		out[v] = struct{}{}
	}
	return out
}

func RawFrom[T comparable](in ...T) map[T]struct{} {
	return RawFromSlice(in)
}

func RawToSlice[T comparable](in map[T]struct{}) []T {
	out := make([]T, 0, len(in))
	for k := range in {
		out = append(out, k)
	}
	return out
}

// s & arr
func RawIntersectSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
	out := make(map[T]struct{})
	for _, v := range arr {
		if _, ok := s[v]; ok {
			out[v] = struct{}{}
		}
	}
	return out
}

// s & other
func RawIntersectSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		if _, ok := other[k]; ok {
			out[k] = struct{}{}
		}
	}
	return out
}

// s - arr
func RawSubtractSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
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
func RawSubtractSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
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
func RawUnionSlice[T comparable](s map[T]struct{}, arr []T) map[T]struct{} {
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
func RawUnionSet[T comparable](s map[T]struct{}, other map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	for k := range other {
		out[k] = struct{}{}
	}
	return out
}

// Merge arr into s, will modify s, return s
func RawMergeSlice[T comparable](s map[T]struct{}, arrs ...[]T) map[T]struct{} {
	for _, a := range arrs {
		for _, v := range a {
			s[v] = struct{}{}
		}
	}
	return s
}

// Merge other into s, will modify s, return s
func RawMergeSet[T comparable](s map[T]struct{}, others ...map[T]struct{}) map[T]struct{} {
	for _, other := range others {
		for k := range other {
			s[k] = struct{}{}
		}
	}
	return s
}

// Remove arr from s, will modify s, return s
func RawRemoveSlice[T comparable](s map[T]struct{}, arrs ...[]T) map[T]struct{} {
	for _, a := range arrs {
		for _, v := range a {
			delete(s, v)
		}
	}
	return s
}

// Remove other from s, will modify s, return s
func RawRemoveSet[T comparable](s map[T]struct{}, others ...map[T]struct{}) map[T]struct{} {
	for _, other := range others {
		for k := range other {
			delete(s, k)
		}
	}
	return s
}

func RawEqual[T comparable](a map[T]struct{}, b map[T]struct{}) bool {
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

func RawCopy[T comparable](s map[T]struct{}) map[T]struct{} {
	out := make(map[T]struct{})
	for k := range s {
		out[k] = struct{}{}
	}
	return out
}

func RawHas[T comparable](s map[T]struct{}, v T) bool {
	_, ok := s[v]
	return ok
}

func RawLen[T comparable](s map[T]struct{}) int      { return len(s) }
func RawIsEmpty[T comparable](s map[T]struct{}) bool { return len(s) == 0 }
func RawClear[T comparable](s map[T]struct{}) {
	for k := range s {
		delete(s, k)
	}
}
func RawAdd[T comparable](s map[T]struct{}, vs ...T) map[T]struct{} { return RawMergeSlice(s, vs) }
func RawDel[T comparable](s map[T]struct{}, vs ...T) map[T]struct{} { return RawRemoveSlice(s, vs) }
