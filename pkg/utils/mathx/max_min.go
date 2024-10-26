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

func Clamp[T constraints.Ordered](in, min, max T) T {
	if in < min {
		return min
	}
	if in > max {
		return max
	}

	return in
}
