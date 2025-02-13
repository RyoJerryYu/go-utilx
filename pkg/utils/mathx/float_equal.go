package mathx

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Float = constraints.Float

// FloatEqualEpsilon is the default threshold for floating point equality comparison.
// Two floating point numbers are considered equal if their difference is less than this value.
const FloatEqualEpsilon = 1e-9

// Equal compares two floating point numbers for equality using the default threshold (FloatEqualEpsilon).
// Returns true if the absolute difference between a and b is less than FloatEqualEpsilon.
//
// Example:
//
//	Equal(1.0, 1.0000000001) // returns true
//	Equal(1.0, 1.1) // returns false
func Equal[T Float](a, b T) bool {
	return EqualWithThreshold(a, b, FloatEqualEpsilon)
}

// EqualWithThreshold compares two floating point numbers for equality using a custom threshold.
// Returns true if the absolute difference between a and b is less than the given threshold.
//
// Example:
//
//	EqualWithThreshold(1.0, 1.01, 0.1) // returns true
//	EqualWithThreshold(1.0, 1.2, 0.1) // returns false
func EqualWithThreshold[T Float](a, b, threshold T) bool {
	return math.Abs(float64(a-b)) < float64(threshold)
}
