package logger_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestEmptyString(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("empty constant should be empty").
		WithActual(logger.EMPTY).
		WithExpected("").
		MustEqual()
}
