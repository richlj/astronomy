// Package astro is used for carrying out astronomical calculations
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

var TestJulianTimeJulianDayData = []struct {
	input  julianTime
	output julianDay
}{
	{input: 24583346.324461, output: 24583346.000000},
	{input: 23437892.876532, output: 23437893.000000},
	{input: 29999999.999999, output: 30000000.000000},
}

func TestJulianTimeJulianDay(t *testing.T) {
	data := TestJulianTimeJulianDayData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.julianDay(); result != output {
			t.Errorf("expected: `%f`; got: `%f`", output,
				result)
		}
	}
}

var TestGregorianTimeJulianData = []struct {
	input  gregorianTime
	output julianTime
}{
	{
		gregorianTime(time.Date(2017, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		2458102.380451,
	},
}

func TestGregorianTimeJulian(t *testing.T) {
	data := TestGregorianTimeJulianData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.julian(); !result.almostEqual(output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

var TestGregorianTimeC19CorrectionData = []struct {
	input  gregorianTime
	output float64
}{
	{
		gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		0.000000,
	},
	{
		gregorianTime(time.Date(1900, 2, 28, 21, 7, 51, 0,
			time.FixedZone("UTC", 0))),
		1.000000,
	},
	{
		gregorianTime(time.Date(1900, 3, 1, 21, 7, 51, 0,
			time.FixedZone("UTC", 0))),
		0.000000,
	},
}

func TestGregorianTimeC19Correction(t *testing.T) {
	data := TestGregorianTimeC19CorrectionData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.c19Correction()
		if !almostEqual(result, output) {
			t.Errorf("expected: `%f`; got: `%f`", output, result)
		}
	}
}

var TestJulianTimeGregorianData = []struct {
	input  julianTime
	output gregorianTime
}{
	{
		julianTime(2460528.38793),
		gregorianTime(time.Date(2024, 8, 5, 21, 18, 37, 0,
			time.FixedZone("UTC", 0))),
	},
	{
		julianTime(2460527.596272),
		gregorianTime(time.Date(2024, 8, 5, 2, 18, 37, 0,
			time.FixedZone("UTC", 0))),
	},
	{
		julianTime(2460619.97127),
		gregorianTime(time.Date(2024, 11, 5, 11, 18, 37, 0,
			time.FixedZone("UTC", 0))),
	},
	{
		julianTime(2451545.13125),
		gregorianTime(time.Date(2000, 1, 1, 15, 9, 0, 0,
			time.FixedZone("UTC", 0))),
	},
	{
		julianTime(2445853.03403),
		gregorianTime(time.Date(1984, 6, 1, 12, 49, 0, 0,
			time.FixedZone("UTC", 0))),
	},
}

func TestJulianTimeGregorian(t *testing.T) {
	data := TestJulianTimeGregorianData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, time.Time(data[i].output)
		result := time.Time(input.gregorian())
		if output.Sub(result) != 0 {
			t.Errorf("expected: `%s`; got: `%s`", result, output)
		}
	}
}

var TestAltitudeCorrectionData = []struct {
	input  Altitude
	output float64
}{
	{input: 0, output: -0.162500},
	{input: 10, output: -0.847523},
	{input: 50, output: -1.895120},
	{input: 100, output: -2.680104},
	{input: 314.159265359, output: -4.750361},
	{input: 1000, output: -8.475235},
}

func TestAltitudeCorrection(t *testing.T) {
	data := TestAltitudeCorrectionData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.correction(); !almostEqual(result, output) {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

type TestLocationHourAngleInput struct {
	location Location
	day      julianDay
}

var TestLocationHourAngleData = []struct {
	input  TestLocationHourAngleInput
	output julianTime
}{
	{
		TestLocationHourAngleInput{
			Location{0, 0, 0}, 12345678,
		},
		91.079161,
	},
	{
		TestLocationHourAngleInput{
			Location{-134.219, 11.462, 0}, 2454449,
		},
		62.219506,
	},
}

func TestLocationHourAngle(t *testing.T) {
	data := TestLocationHourAngleData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.hourAngle(input.day)
		if !result.almostEqual(output) {
			t.Errorf("expected result %f, got result %f", output, result)
		}
	}
}

var TestGregorianTimeJulianDateData = []struct {
	input  gregorianTime
	output julianTime
}{
	{
		gregorianTime(time.Date(2007, 12, 14, 21, 7, 51, 0,
			time.FixedZone("PDT", -25200))),
		2454449.000000,
	},
	{
		gregorianTime(time.Date(2039, 1, 12, 1, 7, 51, 0,
			time.FixedZone("GMT", 0))),
		2465800.000000,
	},
}

func TestGregorianTimeJulianDate(t *testing.T) {
	data := TestGregorianTimeJulianDateData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.julianDate(); !result.almostEqual(output) {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestJulianTimeIsZeroData = []struct {
	input  julianTime
	output bool
}{
	{
		julianTime(math.NaN()),
		true,
	},
	{
		julianTime(J2000Epoch),
		false,
	},
	{
		julianTime(300000),
		false,
	},
	{
		gregorianTime(time.Now()).julian(),
		false,
	},
}

func TestJulianTimeIsZero(t *testing.T) {
	data := TestJulianTimeIsZeroData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.IsZero()
		if output != result {
			t.Errorf("result %t does not match expected output %t",
				result, output)
		}
	}
}

var TestGregorianTimeJulianDayData = []struct {
	input  gregorianTime
	output julianDay
}{
	{
		gregorianTime(time.Date(1980, 1, 1, 1, 1, 1, 1,
			time.FixedZone("UTC", 0))),
		julianDay(2444240.000000),
	},
	{
		gregorianTime(time.Date(2007, 12, 14, 21, 8, 1, 0,
			time.FixedZone("PDT", -25200))),
		julianDay(2454449.000000),
	},
}

func TestGregorianTimeJulianDay(t *testing.T) {
	data := TestGregorianTimeJulianDayData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.julian().julianDay(); result != output {
			t.Errorf("expected result %f, got result %f", output,
				result)
		}
	}
}

var TestGregorianTimeStringData = []struct {
	input  gregorianTime
	output string
}{
	{
		gregorianTime(time.Date(1980, 1, 1, 1, 1, 1, 1,
			time.FixedZone("UTC", 0))),
		"1980-01-01T01:01:01+00:00",
	},
	{
		gregorianTime(time.Date(2007, 12, 14, 21, 8, 1, 0,
			time.FixedZone("PDT", -25200))),
		"2007-12-14T21:08:01-07:00",
	},
	{
		gregorianTime(time.Date(2033, 4, 5, 15, 1, 1, 1,
			time.FixedZone("UTC", +10800))),
		"2033-04-05T15:01:01+03:00",
	},
	{
		gregorianTime(time.Date(1991, 11, 25, 2, 59, 57, 0,
			time.FixedZone("CST", 28800))),
		"1991-11-25T02:59:57+08:00",
	},
	{
		gregorianTime(time.Date(2001, 2, 3, 4, 5, 6, 7,
			time.FixedZone("CST", 28800))),
		"2001-02-03T04:05:06+08:00",
	},
	{
		gregorianTime(time.Time{}),
		"n/a",
	},
}

func TestGregorianTimeString(t *testing.T) {
	data := TestGregorianTimeStringData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		if result := input.String(); output != result {
			t.Errorf("expected: `%s`; got: `%s`", result, output)
		}
	}
}

type sunTimeDataInputs struct {
	location Location
	day      julianDay
}

var TestLocationSunriseTimeData = []struct {
	input  sunTimeDataInputs
	output julianTime
}{
	{
		sunTimeDataInputs{Location{45, 10, 0}, 2500000.5},
		julianTime(2500000.251258),
	},
	{
		sunTimeDataInputs{Location{-60, 35, 0}, 2458397.5},
		julianTime(2458397.3214121),
	},
	{
		sunTimeDataInputs{Location{45, -90, 0}, 2482500.5},
		julianTime(2482500.006067),
	},
}

func TestLocationSunriseTime(t *testing.T) {
	data := TestLocationSunriseTimeData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.sunriseTime(input.day)
		if output.gregorian() != result.gregorian() {
			t.Errorf("expected: `%s`; got: `%s`", result.gregorian(),
				output.gregorian())
		}
	}
}

var TestLocationSunsetTimeData = []struct {
	input  sunTimeDataInputs
	output julianTime
}{
	{
		sunTimeDataInputs{Location{45, 10, 0}, 2500000.5},
		julianTime(2500000.809059),
	},
	{
		sunTimeDataInputs{Location{-60, 35, 0}, 2458397.5},
		julianTime(2458397.884761),
	},
	{
		sunTimeDataInputs{Location{45, -90, 0}, 2482500.5},
		julianTime(2482500.495580),
	},
}

func TestLocationSunsetTime(t *testing.T) {
	data := TestLocationSunsetTimeData
	for i := 0; i < len(data); i++ {
		input, output := data[i].input, data[i].output
		result := input.location.sunsetTime(input.day)
		if output.gregorian() != result.gregorian() {
			t.Errorf("expected: `%s`; got: `%s`", result.gregorian(),
				output.gregorian())
		}
	}
}
