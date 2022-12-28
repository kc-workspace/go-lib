package xerrors_test

import (
	"fmt"
	"testing"

	"github.com/kc-workspace/go-lib/random"
	"github.com/kc-workspace/go-lib/xerrors"
	"github.com/kc-workspace/go-lib/xtests"
)

func Mock() error {
	return String(random.New().FixedAlphaNumericString(10))
}

func String(s string) error {
	return fmt.Errorf(s)
}

func TestErrorHandler(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("empty handler").
		WithActual(xerrors.New()).
		MustNotBeNil()

	assertion.NewName("error length").
		WithExpected(1).
		WithActual(xerrors.New().And(Mock()).And(nil).Length()).
		MustEqual()

	assertion.NewName("and error").
		WithExpected(true).
		WithActual(xerrors.New().And(Mock()).And(nil).HasError()).
		MustEqual()

	var data, err = xerrors.New().AndD("string", Mock())
	assertion.NewName("andD return error").
		WithExpected(true).
		WithActual(err.HasError()).
		MustEqual()
	assertion.NewName("andD return data").
		WithExpected("string").
		WithActual(data).
		MustEqual()

	assertion.NewName("merging").
		WithExpected(2).
		WithActual(xerrors.New().And(Mock()).Merge(xerrors.New().And(Mock())).Length()).
		MustEqual()

	var first = Mock()
	assertion.NewName("get first").
		WithExpected(first).
		WithActual(xerrors.New().And(first).Merge(xerrors.New().And(Mock())).First()).
		MustEqual()

	assertion.NewName("get first nil").
		WithActual(xerrors.New().First()).
		MustBeNil()

	assertion.NewName("string - with error").
		WithExpected(`found '2' errors (total=2)
- a
- b
`).
		WithActual(xerrors.New().And(String("a")).And(String("b")).String()).
		MustEqual()

	assertion.NewName("string - not error").
		WithExpected(`not found any errors`).
		WithActual(xerrors.New().String()).
		MustEqual()

	assertion.NewName("get error - when no error").
		WithActual(xerrors.New().Error()).
		MustBeNil()

	assertion.NewName("get error - when has error").
		WithExpected(`found '1' errors (total=1)
- this is error message
`).
		WithError(xerrors.New().And(fmt.Errorf("this is error message")).Error()).
		MustEqualError()
}
