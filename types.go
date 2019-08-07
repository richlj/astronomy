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

// Location is the three-dimensional position of an object above the globe.
// Latitude and Longitude values are in degrees.
type Location struct {
	Latitude  float64  `json:"latitude" validate:"min=-90,max=90"`
	Longitude float64  `json:"longitude" validate:"min=-180,max=180"`
	Altitude  Altitude `json:"altitude" validate:"min=0"`
}
