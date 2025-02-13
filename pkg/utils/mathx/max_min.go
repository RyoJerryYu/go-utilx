package mathx

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Clamp constrains a value to lie between minimum and maximum values.
// Returns:
// - min if in < min
// - max if in > max
// - in otherwise
//
// Example:
//
//	Clamp(5, 0, 10) // returns 5
//	Clamp(-1, 0, 10) // returns 0
//	Clamp(11, 0, 10) // returns 10
func Clamp[T constraints.Ordered](in, min, max T) T {
	if in < min {
		return min
	}
	if in > max {
		return max
	}

	return in
}
