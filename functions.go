// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"math"
)

const (
	// J2000Epoch is January 1, 2000, 12:00 TT
	J2000Epoch julianTime = 2451545.0
)

// J2000Epoch returns the julianTime of a given julianTime within the standard
// epoch "J2000" in the Julian calendar
func (j julianTime) J2000Epoch() julianTime {
	return j - J2000Epoch + 0.0008
}

// sin provides the Sine of an angle that is provided in degress
func sin(a float64) float64 {
	return math.Sin(a / 180 * math.Pi)
}

// cos provides the Cosine of an angle that is provided in degress
func cos(a float64) float64 {
	return math.Cos(a / 180 * math.Pi)
}

// asin provides the arcsine in degress of the supplied value
func asin(a float64) float64 {
	return math.Asin(a) * 180 / math.Pi
}

// acos provides the arccosine in degress of the supplied value
func acos(a float64) float64 {
	return math.Acos(a) * 180 / math.Pi
}
