package containerx

import "github.com/RyoJerryYu/go-utilx/pkg/container/slicex"

// Any converts a value of type T to interface{} (any).
// This is a helper function for type conversion.
func Any[T any](in T) any {
	return in
}

// ToAny converts a slice of type T to a slice of interface{} (any).
// Example:
//
//	slice := []int{1, 2, 3}
//	result := ToAny(slice)
//	// result = []any{1, 2, 3}
func ToAny[T any](in []T) []any {
	return slicex.To(in, Any)
}
