// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"time"
)

type gregorianTime time.Time

type julianTime float64

// Altitude is the height in meters of an object above sea level
type Altitude float64
