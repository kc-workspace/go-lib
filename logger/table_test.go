package logger_test

import (
	"bytes"
	"testing"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestTableError(t *testing.T) {
	var assertion = xtests.New(t)

	var buffer1 = &bytes.Buffer{}
	var printer = logger.NewPrinter(buffer1)
	var table = logger.NewTable(printer, 2)

	assertion.NewName("error from Row()").
		WithError(table.Row("hello", "world")).
		WithExpected("you never initial table").
		MustContainError()
	assertion.NewName("empty output").
		WithActual(buffer1.String()).
		WithExpected("").
		MustEqual()
	assertion.NewName("silent end").
		WithError(table.End()).
		WithExpected("you never initial table").
		MustContainError()
}

func TestNewTable(t *testing.T) {
	var assertion = xtests.New(t)

	var buffer1 = &bytes.Buffer{}
	var printer = logger.NewPrinter(buffer1)
	var table = logger.NewTable(printer, 99)
	table.SetSize(2).Init()

	assertion.NewName("not error").
		WithError(table.Row("hello", "world")).
		MustNotError()
	assertion.NewName("able to call End()").
		WithError(table.End()).
		MustNotError()
	assertion.NewName("create new row").
		WithActual(buffer1.String()).
		WithExpected(`hello  world
`).
		MustEqual()
}
