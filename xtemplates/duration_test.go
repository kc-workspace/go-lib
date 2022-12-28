package xtemplates_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestDurationFunction(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal duration - hour").
		WithExpected("0.25").
		WithActualAndError(xtemplates.Text("{{ toDuration `15m` `h` }}", mapper.New())).
		MustEqual()

	assertion.NewName("normal duration - minute").
		WithExpected("5").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `m` }}", mapper.New())).
		MustEqual()

	assertion.NewName("normal duration - second").
		WithExpected("300").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `s` }}", mapper.New())).
		MustEqual()

	assertion.NewName("normal duration - millisecond").
		WithExpected("300000").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `ms` }}", mapper.New())).
		MustEqual()

	assertion.NewName("day duration - second").
		WithExpected("86400").
		WithActualAndError(xtemplates.Text("{{ dayToDuration 1 `s` }}", mapper.New())).
		MustEqual()

	assertion.NewName("empty input").
		WithExpected("time: invalid duration \"\"").
		WithActualAndError(xtemplates.Text("{{ toDuration `` `s` }}", mapper.New())).
		Must(xtests.MUST_ERROR, xtests.MUST_CONTAINS_ERROR)

	assertion.NewName("invalid unit").
		WithExpected("cannot convert to unit 'a'").
		WithActualAndError(xtemplates.Text("{{ toDuration `5m` `a` }}", mapper.New())).
		Must(xtests.MUST_ERROR, xtests.MUST_CONTAINS_ERROR)
}
