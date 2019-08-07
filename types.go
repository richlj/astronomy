// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"time"
)

type gregorianTime time.Time

type julianTime float64
