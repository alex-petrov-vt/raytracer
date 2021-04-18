package util

import (
	"math"
)

const float64EqualityThreshold = 1e-7

// AlmostEqual compares two float64 numbers using epsilon
func AlmostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
