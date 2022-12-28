package dotenv_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/dotenv"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestReader(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("read normal environment format").
		WithActualAndError(dotenv.Unmarshal("test=true")).
		WithExpected(map[string]string{
			"test": "true",
		}).
		MustDeepEqual()

	assertion.NewName("read normal environment format").
		WithActualAndError(dotenv.Unmarshal("# test=true")).
		WithExpected(map[string]string{}).
		Must(xtests.MUST_NOT_ERROR, xtests.MUST_DEEP_EQUAL)
}
