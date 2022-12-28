package configs_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/configs"
	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtests"
)

func toMap(key, value string, ok bool) (mapper.Mapper, bool) {
	if ok {
		return mapper.New().Set(key, value), true
	}
	return mapper.New(), false
}

func TestOverrideParser(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("regular string format").
		WithActualAndBool(toMap(configs.ParseOverride("test=true"))).
		WithExpected(mapper.New().Set("test", "true")).
		MustDeepEqual()
	assertion.NewName("with other sign").
		WithActualAndBool(toMap(configs.ParseOverride("%$^=(!)@$"))).
		WithExpected(mapper.New().Set("%$^", "(!)@$")).
		MustDeepEqual()
	assertion.NewName("too many equal sign 1").
		WithActualAndBool(toMap(configs.ParseOverride("a=b=c"))).
		MustError()
	assertion.NewName("too many equal sign 2").
		WithActualAndBool(toMap(configs.ParseOverride("a==b"))).
		MustError()
	assertion.NewName("no equal sign").
		WithActualAndBool(toMap(configs.ParseOverride("abc"))).
		MustError()
}

func TestClusterConfig(t *testing.T) {
	var assertion = xtests.New(t)

	var baseConfig = mapper.New().
		Set("test", true).
		Set("status", "validated").
		Set("_.C1.test", false).
		Set("_.c1.status", "invalid").
		Set("test", true).
		Set("test", true)

	assertion.NewName("get regular config").
		WithActualAndError(configs.BuildClusterConfig("", baseConfig).Get("test")).
		WithExpected(true).
		MustEqual()
	assertion.NewName("get cluster regular config").
		WithActualAndError(configs.BuildClusterConfig("C1", baseConfig).Get("test")).
		WithExpected(false).
		MustEqual()
	assertion.NewName("get cluster regular config (case-insensitive)").
		WithActualAndError(configs.BuildClusterConfig("c1", baseConfig).Get("test")).
		WithExpected(false).
		MustEqual()
}
