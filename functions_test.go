// Package astro contains functions for calculating astronomical times and
// positions
package astro

import (
	"math"
	"testing"
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

func (j julianTime) almostEqual(a julianTime) bool {
	return math.Abs(float64(j)-float64(a)) < tolerance
}
