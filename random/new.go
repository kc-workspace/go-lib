package random

import (
	"math/rand"
	"text/template"
	"time"
)

func New() *Randomer {
	return NewWithSeed(time.Now().UnixNano())
}

func NewWithSeed(seed int64) *Randomer {
	var random = rand.New(rand.NewSource(seed))

	var randomer = &Randomer{
		Seed: time.Now().UnixNano(),

		underlay: random,
		template: template.New("Default"),
	}

	// Add functions to the template
	randomer.template = randomer.template.Funcs(template.FuncMap{
		"bool":      randomer.Bool,
		"int":       randomer.Int32,
		"int64":     randomer.Int64,
		"float":     randomer.Float32,
		"float64":   randomer.Float64,
		"fNumStr":   randomer.FixedNumericString,
		"dNumStr":   randomer.DynamicNumericString,
		"fALowStr":  randomer.FixedAlphabetLowerString,
		"dALowStr":  randomer.DynamicAlphabetLowerString,
		"fAUpStr":   randomer.FixedAlphabetUpperString,
		"dAUpStr":   randomer.DynamicAlphabetUpperString,
		"fAnLowStr": randomer.FixedAlphaNumericLowerString,
		"dAnLowStr": randomer.DynamicAlphaNumericLowerString,
		"fAnUpStr":  randomer.FixedAlphaNumericUpperString,
		"dAnUpStr":  randomer.DynamicAlphaNumericUpperString,
		"fAnStr":    randomer.FixedAlphaNumericString,
		"dAnStr":    randomer.DynamicAlphaNumericString,
		"select":    randomer.Select,
	})

	return randomer
}
