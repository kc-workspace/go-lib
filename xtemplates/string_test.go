package xtemplates_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestJoin(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("joining string").
		WithExpected("a-b-c").
		WithActualAndError(xtemplates.Text(`{{ join "a" "b" "c" }}`, mapper.New())).
		MustEqual()

	assertion.NewName("joining partial empty string").
		WithExpected("a-c").
		WithActualAndError(xtemplates.Text(`{{ join "a" "" "c" }}`, mapper.New())).
		MustEqual()

	assertion.NewName("joining empty string").
		WithExpected("").
		WithActualAndError(xtemplates.Text(`{{ join "" "" "" }}`, mapper.New())).
		MustEqual()

	assertion.NewName("joining empty array").
		WithExpected("").
		WithActualAndError(xtemplates.Text(`{{ join }}`, mapper.New())).
		MustEqual()

	assertion.NewName("joining empty array").
		WithExpected("").
		WithActualAndError(xtemplates.Text(`{{ joinArray .a }}`, mapper.New().Set("a", []interface{}{}))).
		MustEqual()

	assertion.NewName("joining array error when not parameter").
		WithExpected("wrong number of args for joinArray").
		WithActualAndError(xtemplates.Text(`{{ joinArray }}`, mapper.New())).
		MustContainError()

	assertion.NewName("joining array").
		WithExpected("a,b,c").
		WithActualAndError(xtemplates.Text(`{{ joinArray .a }}`, mapper.New().Set("a", []interface{}{"a", "b", "c"}))).
		MustEqual()
}
