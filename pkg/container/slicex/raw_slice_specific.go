package slicex

//////
// Slice Specific
//////

// RemoveAt removes the element at the specified index from the slice.
// The order of remaining elements is preserved.
// Note: This function modifies the original slice.
// Example:
//
//	slice := []int{1, 2, 3, 4}
//	result := RemoveAt(slice, 1)
//	// result = []int{1, 3, 4}
func RemoveAt[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

// Include returns true if sliceX contains all elements from sliceY.
// The order of elements is not considered.
// Example:
//
//	x := []int{1, 2, 3, 4}
//	y := []int{2, 4}
//	Include(x, y) // returns true
//	Include(y, x) // returns false
func Include[T comparable](sliceX []T, sliceY []T) bool {
	set := ToSet(sliceX)
	for _, v := range sliceY {
		if _, ok := set[v]; !ok {
			return false
		}
	}

	return true
}

// ElementEqual returns true if both slices contain the same elements,
// regardless of their order. Each element must appear the same number of times in both slices.
// Example:
//
//	a := []int{1, 2, 2, 3}
//	b := []int{2, 1, 3, 2}
//	ElementEqual(a, b) // returns true
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

// Reverse reverses the order of elements in the slice in place.
// Note: This function modifies the original slice.
// Example:
//
//	slice := []int{1, 2, 3, 4}
//	Reverse(slice)
//	// slice is now []int{4, 3, 2, 1}
func Reverse[T any](in []T) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}

// Chunk splits a slice into smaller slices of the specified size.
// The last chunk may be smaller if the slice length is not divisible by chunkSize.
// Panics if chunkSize is less than 1.
// Example:
//
//	slice := []int{1, 2, 3, 4, 5}
//	result := Chunk(slice, 2)
//	// result = [][]int{{1, 2}, {3, 4}, {5}}
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

// Deduplicate returns a new slice with duplicate elements removed.
// The order of first occurrence of each element is preserved.
// Example:
//
//	slice := []int{1, 1, 2, 3, 1, 2}
//	result := Deduplicate(slice)
//	// result = []int{1, 2, 3}
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

// DeduplicateBack returns a new slice with duplicate elements removed,
// keeping the last occurrence of each element.
// The relative order of remaining elements is preserved.
// Example:
//
//	slice := []int{1, 1, 2, 3, 1, 2}
//	result := DeduplicateBack(slice)
//	// result = []int{3, 1, 2}
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
