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

var (
	earthAngleOfTilt = 23.439281 // Appoximate value

	// unixEpoch is Thursday, 1 January 1970 UTC
	unixEpoch = time.Unix(0, 0).UTC()
)

func (j julianTime) julianDay() julianDay {
	return julianDay(math.Round(float64(j)))
}

// J2000Epoch returns the julianTime of a given julianTime within the standard
// epoch "J2000" in the Julian calendar
func (j julianTime) J2000Epoch() julianTime {
	return j - J2000Epoch + 0.0008
}

// gregorian provides a gregorianTime corresponding to the supplied julianTime
func (j julianTime) gregorian() gregorianTime {
	if t := j - gregorianTime(unixEpoch).julian(); j != 0 {
		return gregorianTime(time.Unix(int64(t*86400), 0))
	}
	return gregorianTime{}
}

// julian converts a gregorianTime into a julianTime (for dates with years in
// the range of 1801 to 2099, inclusive)
func (g gregorianTime) julian() julianTime {
	return julianTime(367*g.year() - float64(int(7*(float64(g.year())+
		float64(int((g.month()+9)/12)))/4)) +
		float64((275*int(g.month()))/9) + 1721013.5 + g.day() +
		g.fractionalDay() + g.c19Correction())
}

// c19Correction supplies a value for correcting the conversion of
// gregorianTimes to julianTimes depending on whether they are after
// 28th February 1900
func (g gregorianTime) c19Correction() float64 {
	if 100*g.year()+g.month()-190002.5 < 0 {
		return 1
	}
	return 0
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

func (a Location) solarDeclination(j julianDay) float64 {
	return asin(sin(a.eclipticLongitude(j)) * sin(earthAngleOfTilt))
}

// Altitude.correction attempts to provide a correcting value for solar
// transit calculations with regards to altitude. It's probably going to be
// quite wide of the mark for values above sea level, and it doesn't even
// attempt to provide accurate value for locations below sea level.
func (a Altitude) correction() float64 {
	if a > 0 {
		return -2.076 * (math.Sqrt(float64(a) / 60))
	}
	return -0.1625
}

func (a Location) hourAngle(j julianDay) julianTime {
	return julianTime(acos((sin(-0.83+a.Altitude.correction()) -
		sin(a.Latitude)*sin(a.solarDeclination(j))) /
		cos(a.Latitude) / cos(a.solarDeclination(j))))
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
