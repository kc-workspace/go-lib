package configs_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/configs"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtests"
)

var log = logger.NewLogger(
	make([]string, 0),
	logger.SILENT,
	logger.DefaultPrinter,
)

// TODO: This test didn't works as I expected, Actually this throw error out because config builder will try to query data from config files/directories as well
func TestCore(t *testing.T) {
	var assertion = xtests.New(t)

	var builder = configs.New("config", mapper.New().Set("test", true), log)
	assertion.NewName("direct configuration").
		WithActualAndError(builder.Build([]string{})).
		WithExpected(mapper.New().Set("test", true)).
		MustDeepEqual()
}
