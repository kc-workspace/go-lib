package xtemplates_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestByCluster(t *testing.T) {
	var assertion = xtests.New(t)
	var config = mapper.New().
		Set("base", "default").
		Set("_.1A.base", "1st ant")

	assertion.NewName("override occurred").
		WithExpected("1st ant").
		WithActualAndError(xtemplates.Text("{{ byCluster . `1A` `base` }}", config)).
		MustEqual()

	assertion.NewName("fallback to default").
		WithExpected("default").
		WithActualAndError(xtemplates.Text("{{ byCluster . `3A` `base` }}", config)).
		MustEqual()
}
