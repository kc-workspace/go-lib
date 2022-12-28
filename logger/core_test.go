package logger_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestGlobalLevel(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("Default global level is INFO").
		WithActual(logger.GetLevel()).
		WithExpected(logger.INFO).
		MustEqual()

	logger.SetLevel(4)
	assertion.NewName("Able to setup global level via SetLevel").
		WithActual(logger.GetLevel()).
		WithExpected(logger.DEBUG).
		MustEqual()

	logger.SetLevel(-72)
	assertion.NewName("SetLevel should round if user enter wrong level").
		WithActual(logger.GetLevel()).
		WithExpected(logger.SILENT).
		MustEqual()
}
