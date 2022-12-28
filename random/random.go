package random

import (
	"bytes"
	"math/rand"
	"text/template"
	"time"
)

type Randomer struct {
	Seed int64

	underlay *rand.Rand
	template *template.Template
}

func (randomer *Randomer) UUID() string {
	return randomer.Template("{{ fAnLStr 8 }}-{{ fAnLStr 4 }}-{{ fAnLStr 4 }}-{{ fAnLStr 4 }}-{{ fAnLStr 12 }}")
}

func (randomer *Randomer) IPAddress() string {
	return randomer.Template("{{ int 0 256 }}.{{ int 0 256 }}.{{ int 0 256 }}.{{ int 0 256 }}")
}

func (randomer *Randomer) Bool() bool {
	return randomer.underlay.Intn(2) == 0
}

func (randomer *Randomer) Int32(minimum, maximum int) int {
	return randomer.underlay.Intn(maximum-minimum) + minimum
}

func (randomer *Randomer) Int64(minimum, maximum int64) int64 {
	return randomer.underlay.Int63n(maximum-minimum) + minimum
}

func (randomer *Randomer) Float32(minimum, maximum float32) float32 {
	return randomer.underlay.Float32()*(maximum-minimum) + minimum
}

func (randomer *Randomer) Float64(minimum, maximum float64) float64 {
	return randomer.underlay.Float64()*(maximum-minimum) + minimum
}

func (randomer *Randomer) FixedNumericString(length int) string {
	return randomer.FixedString(length, Numeric)
}

func (randomer *Randomer) DynamicNumericString(minimum, maximum int) string {
	return randomer.DynamicString(minimum, maximum, Numeric)
}

func (randomer *Randomer) FixedAlphabetLowerString(length int) string {
	return randomer.FixedString(length, AlphabetLower)
}

func (randomer *Randomer) DynamicAlphabetLowerString(minimum, maximum int) string {
	return randomer.DynamicString(minimum, maximum, AlphabetLower)
}

func (randomer *Randomer) FixedAlphabetUpperString(length int) string {
	return randomer.FixedString(length, AlphabetUpper)
}

func (randomer *Randomer) DynamicAlphabetUpperString(minimum, maximum int) string {
	return randomer.DynamicString(minimum, maximum, AlphabetUpper)
}

func (randomer *Randomer) FixedAlphaNumericLowerString(length int) string {
	return randomer.FixedString(length, AlphaNumericLower)
}

func (randomer *Randomer) DynamicAlphaNumericLowerString(minimum, maximum int) string {
	return randomer.DynamicString(minimum, maximum, AlphaNumericLower)
}

func (randomer *Randomer) FixedAlphaNumericUpperString(length int) string {
	return randomer.FixedString(length, AlphaNumericUpper)
}

func (randomer *Randomer) DynamicAlphaNumericUpperString(minimum, maximum int) string {
	return randomer.DynamicString(minimum, maximum, AlphaNumericUpper)
}

func (randomer *Randomer) FixedAlphaNumericString(length int) string {
	return randomer.FixedString(length, AlphaNumeric)
}

func (randomer *Randomer) DynamicAlphaNumericString(minimum, maximum int) string {
	return randomer.DynamicString(minimum, maximum, AlphaNumeric)
}

func (randomer *Randomer) FixedString(length int, mode StringMode) string {
	allowance := selectRune(mode)
	s := make([]rune, length)
	for i := range s {
		s[i] = allowance[rand.Intn(len(allowance))]
	}
	return string(s)
}

func (randomer *Randomer) DynamicString(minimum, maximum int, mode StringMode) string {
	var length = randomer.Int32(minimum, maximum)
	return randomer.FixedString(length, mode)
}

func (randomer *Randomer) Utc() time.Time {
	var year = randomer.Int32(2000, 2041)
	var day = randomer.Int32(1, 366)
	var hour = randomer.Int32(0, 24)
	var min = randomer.Int32(0, 61)
	var sec = randomer.Int32(0, 61)
	var nsec = randomer.Int32(0, 1e9)

	return time.Date(year, 1, day, hour, min, sec, nsec, time.UTC)
}

// ShiftNow is same with ShiftDate but using Now as date
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func (randomer *Randomer) ShiftNow(s string) time.Time {
	return randomer.ShiftDate(s, time.Now())
}

// ShiftDate will return future/past date base on given s as maximum shift
// For shift to past use negative syntax
// Example s string: 1d, 24h, -360m, 30s.
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
func (randomer *Randomer) ShiftDate(s string, date time.Time) time.Time {
	dur, err := time.ParseDuration(s)
	if err != nil {
		panic(err)
	}

	var randomNumber = randomer.Int64(1, dur.Nanoseconds())
	return time.Now().Add(time.Duration(randomNumber))
}

func (randomer *Randomer) Select(values ...interface{}) interface{} {
	var index = randomer.Int32(0, len(values))
	return values[index]
}

type WeightedValue struct {
	Weight int
	Value  interface{}
}

// WeightedSelect will return a random value from the given values with the given weights.
// This will run O(2n)
func (randomer *Randomer) WeightedSelect(values ...WeightedValue) interface{} {
	var totalWeight int
	for i := 0; i < len(values); i++ {
		var weight = values[i].Weight
		if weight <= 0 {
			panic("Weight must be greater than 0")
		}

		totalWeight += weight
	}

	var chance = randomer.Int32(0, totalWeight)
	for i := 0; i < len(values); i++ {
		var value = values[i]
		if chance < value.Weight {
			return value.Value
		}
		chance -= value.Weight
	}

	return nil
}

func (randomer *Randomer) Template(tmpl string) string {
	var parsed, err = randomer.template.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	err = parsed.Execute(&buffer, nil)
	if err != nil {
		panic(err)
	}

	return buffer.String()
}
