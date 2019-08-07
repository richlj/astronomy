// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"math"
	"time"

	"github.com/go-validator/validator"
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

// meanSolarNoon provides the Julian 2000 Epoch julianTime of the mean solar
// noon for a given Location on a particlular julianDay
func (a Location) meanSolarNoon(j julianDay) julianTime {
	return julianTime(j).J2000Epoch() + julianTime(a.Longitude/360)
}

func (a Location) solarMeanAnomaly(j julianDay) float64 {
	return math.Mod(357.5291+0.98560028*float64(a.meanSolarNoon(j)),
		360)
}

func (a Location) equationOfTheCentre(j julianDay) float64 {
	sma := a.solarMeanAnomaly(j)
	return 1.9148*math.Sin(a.Longitude) + 0.0200*math.Sin(2*sma) +
		0.0003*math.Sin(3*sma)
}

func (a Location) eclipticLongitude(j julianDay) float64 {
	return math.Mod(a.solarMeanAnomaly(j)+a.equationOfTheCentre(j)+
		180+102.9732, 360)
}

func (a Location) solarTransit(j julianDay) julianTime {
	return J2000Epoch + a.meanSolarNoon(j) +
		julianTime(0.0053*math.Sin(a.solarMeanAnomaly(j)-
			0.0069*sin(2*a.eclipticLongitude(j))))
}

func (g gregorianTime) fractionalDay() float64 {
	return (g.hour()*3600 + g.minute()*60 + g.second()) / 86400
}

func (g gregorianTime) year() float64 {
	return float64(time.Time(g).Year())
}

func (g gregorianTime) month() float64 {
	return float64(time.Time(g).Month())
}

func (g gregorianTime) day() float64 {
	return float64(time.Time(g).Day())
}

func (g gregorianTime) hour() float64 {
	return float64(time.Time(g).Hour())
}

func (g gregorianTime) minute() float64 {
	return float64(time.Time(g).Minute())
}

func (g gregorianTime) second() float64 {
	return float64(time.Time(g).Second())
}

func (g gregorianTime) date() (int, int, int) {
	y, m, d := time.Time(g).Date()
	return y, int(m), d
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

func (a Location) validate() error {
	return validator.Validate(a)
}
