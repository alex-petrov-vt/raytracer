package util

import (
	"math"
)

const float64EqualityThreshold = 1e-5

// FloatEquals compares two float64 numbers using epsilon
func FloatEquals(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
