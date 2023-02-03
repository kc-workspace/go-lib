package configs_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/configs"
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestParseConfigFromEnv(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("normal environment").
		WithActualAndError(configs.ParseConfigFromEnv("FTH", []string{"FTH_FREQTRADE_STATUS=running"})).
		WithExpected(mapper.New().Set("freqtrade-status", "running")).
		MustDeepEqual()
	assertion.NewName("unknown environment").
		WithActualAndError(configs.ParseConfigFromEnv("FTH", []string{"FTHSTATUS=running"})).
		WithExpected(mapper.New()).
		MustDeepEqual()
}
