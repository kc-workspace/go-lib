package xtemplates_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtemplates"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestXtemplate(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("invalid template").
		WithExpected("function \"invalid\" not defined").
		WithActualAndError(xtemplates.Text("{{ invalid \"function\" }}", mapper.New())).
		MustContainError()
}
