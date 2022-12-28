package logger_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestToLevel(t *testing.T) {
	var assertion = xtests.New(t)
	for _, tc := range []xtests.TestCase{
		xtests.NewCase("silent level", float64(0), logger.SILENT),
		xtests.NewCase("error level", float64(1), logger.ERROR),
		xtests.NewCase("warn level", float64(2), logger.WARN),
		xtests.NewCase("info level", float64(3), logger.INFO),
		xtests.NewCase("debug level", float64(4), logger.DEBUG),
		xtests.NewCase("negative level", float64(-20), logger.SILENT),
		xtests.NewCase("exceeded level", float64(99), logger.DEBUG),
	} {
		assertion.NewName(tc.Name).
			WithActual(logger.ToLevel(tc.Actual.(float64))).
			WithExpected(tc.Expected).
			MustEqual()
	}
}
