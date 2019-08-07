// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"math"
)

// tolerance is used for comparing float values in tests
var tolerance = math.Pow(10, -6)

// almostEqual compares the values of two float64s within a set parameter
func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= tolerance
}
