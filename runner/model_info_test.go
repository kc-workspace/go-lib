package runner_test

import (
	"errors"
	"testing"
	"time"

	"github.com/kc-workspace/go-lib/runner"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestBasicInformation(t *testing.T) {
	var assertion = xtests.New(t)

	var info = runner.NewInformation("test")
	assertion.NewName("default name").
		WithExpected("test").
		WithActual(info.Name()).
		MustEqual()
	assertion.NewName("default status").
		WithExpected(runner.INITIAL).
		WithActual(info.Status()).
		MustEqual()
	assertion.NewName("default duration").
		WithExpected(time.Duration(-1)).
		WithActual(info.Duration()).
		MustEqual()
	assertion.NewName("isInitial return true").
		WithExpected(true).
		WithActual(info.IsInitial()).
		MustEqual()
	assertion.NewName("toString").
		WithExpected("test: initial (-1ns)").
		WithActual(info.String()).
		MustEqual()
}

func TestSetterInformation(t *testing.T) {
	var assertion = xtests.New(t)

	var info = runner.NewInformation("test")
	assertion.NewName("able to set duration").
		WithExpected(5 * time.Second).
		WithActual(info.SetDuration(5 * time.Second).Duration()).
		MustEqual()
	assertion.NewName("able to set error message").
		WithExpected("test").
		WithActual(info.SetError(errors.New("test")).Error().Error()).
		MustEqual()
	assertion.NewName("able to set status").
		WithExpected(runner.INVALID).
		WithActual(info.SetStatus(runner.INVALID).Status()).
		MustEqual()
	assertion.NewName("set status works only once").
		WithExpected(runner.INVALID).
		WithActual(info.SetStatus(runner.ERROR).Status()).
		MustEqual()
}
