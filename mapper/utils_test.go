package mapper_test

import (
	"testing"

	"github.com/kc-workspace/go-lib/mapper"
	"github.com/kc-workspace/go-lib/xtests"
)

func TestGetUtilities(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("valid data").
		WithExpected("world").
		WithActualAndError(mapper.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": "world",
			},
		}, "test.hello")).
		MustEqual()

	assertion.NewName("valid object data").
		WithExpected(map[string]interface{}{
			"hello": "world",
		}).
		WithActualAndError(mapper.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": "world",
			},
		}, "test")).
		MustDeepEqual()

	assertion.NewName("missing data").
		WithExpected(nil).
		WithActualAndError(mapper.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": "world",
			},
		}, "test.test")).
		MustEqual()

	assertion.NewName("nil data").
		WithExpected(nil).
		WithActualAndError(mapper.Get(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": nil,
			},
		}, "test.hello")).
		MustEqual()
}

func TestGetsUtilities(t *testing.T) {
	var assertion = xtests.New(t)

	assertion.NewName("valid data").
		WithExpected(map[string]interface{}{
			"testing": true,
		}).
		WithActualAndError(mapper.Gets(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": nil,
				"hello2": map[string]interface{}{
					"testing": true,
				},
			},
		}, "test.hello", "test.hello2")).
		MustDeepEqual()

	assertion.NewName("missing data").
		WithExpected(nil).
		WithActualAndError(mapper.Gets(map[string]interface{}{
			"test": map[string]interface{}{
				"hello": nil,
				"hello2": map[string]interface{}{
					"testing": true,
				},
			},
		}, "test.hello", "test.hello3", "test.hello4")).
		MustDeepEqual()
}

func TestNormalize(t *testing.T) {
	var i = mapper.New().
		Set("$schema", "hello").
		Set("internal.test", true).
		Set("internal.#comment#", "hello").
		Set("b.test", "$schema").
		Set("b.#comment#", mapper.New().Set("test", 123)).
		Set("#comment#", mapper.New().Set("test", 123))

	var n = mapper.Normalize(i.Copy(), []string{"$schema", "#comment#"})

	t.Log(i)
	t.Log(n)
}
