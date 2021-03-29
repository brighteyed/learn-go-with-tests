package util

import "math"

const Float64EqualityThreshold = 1e-9

// RoughlyEqualFloat64 tests to values for equality with given threshold
func RoughlyEqualFloat64(a, b float64, threshold float64) bool {
	return math.Abs(a-b) < threshold
}
