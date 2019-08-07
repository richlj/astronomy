// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"fmt"
	"math"
	"testing"
	"time"
)

// tolerance is used for comparing float values in tests
var tolerance = math.Pow(10, -6)

// almostEqual compares the values of two float64s within a set parameter
func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= tolerance
}

var TestSinData = []struct {
	input  float64
	output float64
}{
	{input: 0, output: 0},
	{input: 30, output: 0.5},
	{input: 90, output: 1},
}

func TestSin(t *testing.T) {
	data := TestSinData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := sin(input); !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output,
				result)
		}
	}
}

var TestCosData = []struct {
	input  float64
	output float64
}{
	{input: 0, output: 1},
	{input: 30, output: 0.866025},
	{input: 60, output: 0.5},
	{input: 90, output: 0},
}

func TestCos(t *testing.T) {
	data := TestCosData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := cos(input); !almostEqual(result, output) {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestASinData = []struct {
	input  float64
	output float64
}{
	{input: 0, output: 0},
	{input: 0.5, output: 30},
	{input: 1, output: 90},
}

func TestASin(t *testing.T) {
	data := TestASinData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := asin(input); !almostEqual(result, output) {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestACosData = []struct {
	input  float64
	output float64
}{
	{input: 0, output: 90},
	{input: 0.5, output: 60},
	{input: 1, output: 0},
}

func TestACos(t *testing.T) {
	data := TestACosData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := acos(input); !almostEqual(result, output) {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestJulianTimeJ2000EpochData = []struct {
	input  julianTime
	output julianTime
}{
	{input: 24583346.324461, output: 22131801.325261},
	{input: 23437892.876532, output: 20986347.877332},
	{input: 29999999.999999, output: 27548455.000799},
}

func TestJulianTimeJ2000Epoch(t *testing.T) {
	data := TestJulianTimeJ2000EpochData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.J2000Epoch(); !result.almostEqual(output) {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

type TestLocationMeanSolarNoonInput struct {
	location Location
	day      julianDay
}

var TestLocationMeanSolarNoonData = []struct {
	input  TestLocationMeanSolarNoonInput
	output julianTime
}{
	{
		TestLocationMeanSolarNoonInput{
			Location{0, 0, 0}, 2453954,
		},
		2409.000800,
	},
	{
		TestLocationMeanSolarNoonInput{
			Location{51.5, -0.12462, 0}, 2464546,
		},
		13001.000454,
	},
}

func TestLocationMeanSolarNoon(t *testing.T) {
	data := TestLocationMeanSolarNoonData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.meanSolarNoon(input.day)
		if !result.almostEqual(output) {
			t.Errorf("expected `%f`; got `%f`", output, result)
		}
	}
}

type TestLocationSolarMeanAnomalyInput struct {
	location Location
	day      julianDay
}

var TestLocationSolarMeanAnomalyData = []struct {
	input  TestLocationSolarMeanAnomalyInput
	output float64
}{
	{
		TestLocationSolarMeanAnomalyInput{
			Location{0, 0, 0}, 23437892.000000,
		},
		347.009266,
	},
	{
		TestLocationSolarMeanAnomalyInput{
			Location{32, -120, 0}, 23437892.000000,
		},
		346.680732,
	},
}

func TestLocationSolarMeanAnomaly(t *testing.T) {
	data := TestLocationSolarMeanAnomalyData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.solarMeanAnomaly(input.day)
		if !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

type TestLocationEquationOfTheCentreInput struct {
	location Location
	day      julianDay
}

var TestLocationEquationOfTheCentreData = []struct {
	input  TestLocationEquationOfTheCentreInput
	output float64
}{
	{
		TestLocationEquationOfTheCentreInput{
			Location{0, 0, 0}, 23437892.000000,
		},
		0.005126,
	},
	{
		TestLocationEquationOfTheCentreInput{
			Location{-43.1415, 112.23626, 0}, 2454192.000000,
		},
		-1.464470,
	},
}

func TestLocationEquationOfTheCentre(t *testing.T) {
	data := TestLocationEquationOfTheCentreData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.equationOfTheCentre(input.day)
		if !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

type TestLocationEclipticLongitudeInput struct {
	location Location
	day      julianDay
}

var TestLocationEclipticLongitudeData = []struct {
	input  TestLocationEclipticLongitudeInput
	output float64
}{
	{
		TestLocationEclipticLongitudeInput{
			Location{0, 0, 0}, 0,
		},
		-2.936267,
	},
	{
		TestLocationEclipticLongitudeInput{
			Location{34.2, 11.2, 0}, 22131859,
		},
		41.662002,
	},
}

func TestLocationEclipticLongitude(t *testing.T) {
	data := TestLocationEclipticLongitudeData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.eclipticLongitude(input.day)
		if !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

type LocationSolarTransitInput struct {
	location Location
	day      julianDay
}

var TestLocationSolarTransitData = []struct {
	input  LocationSolarTransitInput
	output julianTime
}{
	{
		LocationSolarTransitInput{
			Location{0, 0, 0}, 12345678,
		},
		12345677.995510,
	},
	{
		LocationSolarTransitInput{
			Location{34.219, 11.462, 0}, 2454449,
		},
		2454449.034946,
	},
}

func TestLocationSolarTransit(t *testing.T) {
	data := TestLocationSolarTransitData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.solarTransit(input.day)
		if !result.almostEqual(output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

var TestGregorianTimeFractionalDayData = []struct {
	input  gregorianTime
	output float64
}{
	{
		gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		0.880451,
	},
}

func TestGregorianTimeFractionalDay(t *testing.T) {
	data := TestGregorianTimeFractionalDayData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.fractionalDay()
		if !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

var TestGregorianTimeYearData = []struct {
	input  gregorianTime
	output float64
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: 2007,
	},
}

func TestGregorianTimeYear(t *testing.T) {
	data := TestGregorianTimeYearData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.year(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeMonthData = []struct {
	input  gregorianTime
	output float64
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: 12,
	},
}

func TestGregorianTimeMonth(t *testing.T) {
	data := TestGregorianTimeMonthData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.month(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeDayData = []struct {
	input  gregorianTime
	output float64
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: 14,
	},
}

func TestGregorianTimeDay(t *testing.T) {
	data := TestGregorianTimeDayData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.day(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeHourData = []struct {
	input  gregorianTime
	output float64
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: 21,
	},
}

func TestGregorianTimeHour(t *testing.T) {
	data := TestGregorianTimeHourData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.hour(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeMinuteData = []struct {
	input  gregorianTime
	output float64
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: 7,
	},
}

func TestGregorianTimeMinute(t *testing.T) {
	data := TestGregorianTimeMinuteData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.minute(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeSecondData = []struct {
	input  gregorianTime
	output float64
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: 51,
	},
}

func TestGregorianTimeSecond(t *testing.T) {
	data := TestGregorianTimeSecondData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.second(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeDateData = []struct {
	input  gregorianTime
	output []int
}{
	{
		input: gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		output: []int{2007, 12, 14},
	},
}

func TestGregorianTimeDate(t *testing.T) {
	data := TestGregorianTimeDateData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		y, m, d := input.date()
		if y != output[0] || m != output[1] || d != output[2] {
			t.Errorf("expected result %d, got result %d %d %d",
				output, y, m, d)
		}
	}
}

func (j julianTime) almostEqual(a julianTime) bool {
	return math.Abs(float64(j)-float64(a)) < tolerance
}

var TestLocationValidateData = []struct {
	input  Location
	output error
}{
	{
		input:  Location{-56.3762, +181.26, 0},
		output: fmt.Errorf("Longitude: greater than max"),
	},
	{
		input:  Location{+106.327, -48.5672, 0},
		output: fmt.Errorf("Latitude: greater than max"),
	},
	{
		input:  Location{+36.3737, +25.373181, 0},
		output: nil,
	},
}

func TestLocationValidate(t *testing.T) {
	data := TestLocationValidateData
	for i := 0; i < len(data); i++ {
		input, out := data[i].input, data[i].output
		result := input.validate()
		if result != nil && out != nil && result.Error() != out.Error() ||
			result != nil && out == nil || result == nil && out != nil {
			t.Errorf("expected `%s`; got: `%s`", out, result)
		}
	}
}

type LocationSolarDeclinationInput struct {
	location Location
	day      julianDay
}

var TestLocationSolarDeclinationData = []struct {
	input  LocationSolarDeclinationInput
	output float64
}{
	{
		LocationSolarDeclinationInput{
			Location{0, 0, 0}, 12345678,
		},
		-23.117070,
	},
	{
		LocationSolarDeclinationInput{
			Location{-134.219, 11.462, 0}, 2454449,
		},
		-23.135386,
	},
}

func TestLocationSolarDeclination(t *testing.T) {
	data := TestLocationSolarDeclinationData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.solarDeclination(input.day)
		if !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}
