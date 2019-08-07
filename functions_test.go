// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"math"
)

// tolerance is used for comparing float values in tests
var tolerance = math.Pow(10, -6)
