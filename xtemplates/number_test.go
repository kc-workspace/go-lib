package xtemplates_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestRatio(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("ratio - 1%").
		WithExpected(`0.01`).
		WithActualAndError(xtemplates.Text(`{{ ratio "1%" }}`, mapper.New())).
		MustEqual()
	assertion.NewName("ratio - 101%").
		WithExpected(`1.01`).
		WithActualAndError(xtemplates.Text(`{{ ratio "101%" }}`, mapper.New())).
		MustEqual()
	assertion.NewName("ratio - 100").
		WithExpected(`1`).
		WithActualAndError(xtemplates.Text(`{{ ratio "100" }}`, mapper.New())).
		MustEqual()
	assertion.NewName("wrong ratio").
		WithActualAndError(xtemplates.Text(`{{ ratio "100a" }}`, mapper.New())).
		MustError()
}
