package mathx

import (
	"math"
)

// NormalizationLog normalizes the data to [0, 1] by log
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

// Normalization normalizes the data to [0, 1]
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
