package mapper_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestNewMapper(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("new mapper").
		WithActual(mapper.New()).
		MustNotBeNil()
	assertion.NewName("new mapper length").
		WithExpected(0).
		WithActual(len(mapper.New())).
		MustEqual()
	assertion.NewName("new mapper is empty").
		WithExpected(true).
		WithActual(mapper.New().IsEmpty()).
		MustEqual()
}

func TestSetMapper(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("set mapper").
		WithExpected(mapper.Mapper(map[string]interface{}{
			"a": true,
		})).
		WithActual(mapper.New().Set("a", true)).
		MustDeepEqual()

	assertion.NewName("set recursive mapper").
		WithExpected(mapper.Mapper(map[string]interface{}{
			"a": map[string]interface{}{
				"b": true,
			},
		})).
		WithActual(mapper.New().Set("a.b", true)).
		MustDeepEqual()
}

func TestGetMapper(t *testing.T) {
	var assertion = xtests.New(t)

	var mapper = mapper.New().
		Set("a", "ant").
		Set("b", 199).
		Set("c", true).
		Set("d", 0.002).
		Set("e", mapper.New().
			Set("ea", "eat").
			Set("eb", 299).
			Set("ec", false).
			Set("ed", []string{"a", "b", "c"})).
		Set("f", []int{9, 8, 7})

	assertion.NewName("get root value").
		WithExpected(int64(199)).
		WithActualAndError(mapper.Ie("b")).
		MustEqual()

	assertion.NewName("get array value").
		WithExpected([]interface{}{9, 8, 7}).
		WithActualAndError(mapper.Ae("f")).
		MustDeepEqual()

	assertion.NewName("error on not exist value").
		WithActualAndError(mapper.Ze("not-exist-value")).
		MustError()
}
