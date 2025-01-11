package slicex

//////
// Slice Specific
//////

func RemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// X > Y
func Include[T comparable](sliceX []T, sliceY []T) bool {
	set := ToSet(sliceX)
	for _, v := range sliceY {
		if _, ok := set[v]; !ok {
			return false
		}
	}

	return true
}

// ElementEqual checks if two slices have the same elements, but not necessarily in the same order
func ElementEqual[T comparable](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	showCnt := map[T]uint64{}
	for _, ele := range a {
		showCnt[ele]++
	}

	for _, ele := range b {
		if showCnt[ele] <= 0 {
			return false
		}
		showCnt[ele]--
	}

	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// reverse the slice in place
func Reverse[T any](in []T) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}

func Chunk[T any](in []T, chunkSize int) [][]T {
	if chunkSize < 1 {
		panic("SliceChunk chunkSize is less than 1")
	}
	out := [][]T{}
	for i := 0; i < len(in); i += chunkSize {
		out = append(out, in[i:min(i+chunkSize, len(in))])
	}

	return out
}

// [1,1,2,3,1,2] -> [1,2,3]
func Deduplicate[T comparable](in []T) []T {
	out := make([]T, 0, len(in))
	seen := make(map[T]struct{})
	for _, v := range in {
		if _, ok := seen[v]; !ok {
			out = append(out, v)
			seen[v] = struct{}{}
		}
	}
	return out
}

// [1,1,2,3,1,2] -> [3,1,2]
func DeduplicateBack[T comparable](in []T) []T {
	out := make([]T, 0, len(in))
	seen := make(map[T]struct{})
	l := len(in) - 1
	for l >= 0 {
		v := in[l]
		if _, ok := seen[v]; !ok {
			out = append(out, v)
			seen[v] = struct{}{}
		}
		l--
	}
	Reverse(out)
	return out
}
