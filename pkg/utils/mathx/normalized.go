package mathx

import (
	"math"
)

// NormalizationLog normalizes the data to [0, 1] using logarithmic scaling.
// The function applies the following steps:
// 1. Applies log(x + 1) to each value
// 2. Scales the logged values to [0, 1] range
//
// This is useful for data with large variations in magnitude.
// All input values must be non-negative.
//
// Example:
//
//	input := []float64{0.15, 0.3, 0.8, 0.5, 2.4, 3.6}
//	output := NormalizationLog(input)
//	// output ≈ [0, 0.088, 0.323, 0.192, 0.782, 1]
func NormalizationLog(data []float64) []float64 {
	out := make([]float64, len(data))
	max, min := float64(-1), math.MaxFloat64

	for idx := range data {
		out[idx] = math.Log(float64(data[idx] + 1))
		max = Max(max, out[idx])
		min = Min(min, out[idx])
	}

	detal := max - min
	for idx := range out {
		out[idx] = (out[idx] - min) / detal
	}

	return out
}

// Normalization scales the input data to the range [0, 1] using min-max normalization.
// The formula used is: (x - min) / (max - min)
//
// Example:
//
//	input := []float64{0.15, 0.3, 0.8, 0.5, 2.4, 3.6}
//	output := Normalization(input)
//	// output ≈ [0, 0.043, 0.188, 0.101, 0.652, 1]
func Normalization(data []float64) []float64 {
	out := make([]float64, len(data))
	max, min := float64(-1), math.MaxFloat64

	for idx := range data {
		max = Max(max, data[idx])
		min = Min(min, data[idx])
	}

	detal := max - min
	for idx := range out {
		out[idx] = (data[idx] - min) / detal
	}

	return out
}
