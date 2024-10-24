package mathx

import (
	"math"

	"golang.org/x/exp/constraints"
)

type Float = constraints.Float

const FloatEqualEpsilon = 1e-9

func Equal[T Float](a, b T) bool {
	return EqualWithThreshold(a, b, FloatEqualEpsilon)
}

func EqualWithThreshold[T Float](a, b, threshold T) bool {
	return math.Abs(float64(a-b)) < float64(threshold)
}
